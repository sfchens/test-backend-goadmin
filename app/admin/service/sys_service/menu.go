package sys_service

import (
	"csf/app/admin/model/sys_model"
	"csf/app/admin/request/sys_request"
	"csf/common/mysql/model"
	"csf/library/easy_db"
	"csf/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"strings"
)

type sSysMenuService struct {
	ctx *gin.Context
}

func NewSysMenuService(ctx *gin.Context) *sSysMenuService {
	return &sSysMenuService{
		ctx: ctx,
	}
}

func (s *sSysMenuService) TreeList(input sys_request.MenuListReq) (out sys_request.MenuListRes, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize

		sysMenuListTmp []sys_model.MenuListItem
	)
	model := s.GetQuery(input)

	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model.Offset((page - 1) * pageSize).Preload("Children").Limit(pageSize).Scan(&sysMenuListTmp).Error
	if err != nil {
		return
	}

	//out.List = s.GetMenuItem(sysMenuListTmp, true)
	list := s.GetMenuItem(sysMenuListTmp, true)
	out.List = s.DealTreeList(list)
	return
}

func (s *sSysMenuService) TreeListAll(input sys_request.MenuListReq) (out sys_request.MenuListRes, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize

		sysMenuListTmp []sys_model.MenuListItem
	)
	model := s.GetQuery(input)

	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model.Offset((page - 1) * pageSize).Preload("Children").Limit(pageSize).Order("sort ASC").Scan(&sysMenuListTmp).Error
	if err != nil {
		return
	}

	list := s.GetMenuItem(sysMenuListTmp, false)
	out.List = s.DealTreeList(list)
	return
}

func (s *sSysMenuService) DealTreeList(data []sys_model.MenuListItem) (res []sys_model.SysMenuListItem) {
	for _, item := range data {
		var tmpV sys_model.SysMenuListItem
		utils.StructToStruct(item.SysMenu, &tmpV)
		tmpV.Label = item.Title
		// 取出apis
		if item.ApisId != "" {
			var apiIds []int
			for _, apiId := range strings.Split(item.ApisId, ",") {
				tmp, _ := strconv.Atoi(apiId)
				apiIds = append(apiIds, tmp)
			}
			tmpV.ApisId = apiIds
			var apiList []model.SysApi
			easy_db.GetDb().Model(model.SysApi{}).Where(fmt.Sprintf("id in (%v)", item.ApisId)).Scan(&apiList)
			tmpV.SysApis = apiList
		} else {
			out, _ := s.GetApisByMenuId([]int{item.Id})
			tmpV.ApisId = out.ApisId
			tmpV.SysApis = out.ApisList
		}
		if len(item.Children) > 0 {
			tmpV.Children = s.DealTreeList(item.Children)
		}

		if len(tmpV.ApisId) <= 0 {
			tmpV.ApisId = []int{}
		}
		res = append(res, tmpV)
	}
	return
}

func (s *sSysMenuService) GetApisByMenuId(ids []int) (out sys_model.GetApisByMenuIdOut, err error) {
	idsArr := s.IntToStringArray(ids)
	idsStr := strings.Join(idsArr, ",")

	err = easy_db.GetDb().Model(model.SysApi{}).Where(fmt.Sprintf("id in (%v)", idsStr)).Scan(&out.ApisList).Error
	if err != nil {
		return
	}

	for _, item := range out.ApisList {
		out.ApisId = append(out.ApisId, int(item.ID))
	}
	return
}

func (s *sSysMenuService) GetQuery(input sys_request.MenuListReq) *gorm.DB {
	var (
		title  = input.Key
		isShow = input.IsShow

		sysMenu model.SysMenu
	)

	model := easy_db.GetDb().Model(sysMenu)

	if title != "" {
		model.Where("title like '%?%'", title)
	}

	if isShow != -1 {
		model.Where("is_show = ?", isShow)
	}
	//model.Where("menu_type=?", "C")
	model.Where("id=?", 2)
	return model
}

func (s *sSysMenuService) GetMenuItem(list []sys_model.MenuListItem, isAll bool) (res []sys_model.MenuListItem) {
	for _, v := range list {
		model1 := easy_db.GetDb().Model(model.SysMenu{}).Preload("Children").Where("menu_type != ?", "F").Where("parent_id = ?", v.Id)
		model1.Order("sort ASC").Scan(&v.Children)

		if len(v.Children) > 0 {
			v.Children = s.GetMenuItem(v.Children, isAll)
		}
		if !isAll {
			var apis []sys_model.MenuListItem
			easy_db.GetDb().Model(model.SysMenu{}).Preload("Children").
				Where("menu_type = ?", "F").
				Where("parent_id = ?", v.Id).Order("sort ASC").Scan(&apis)
			if len(apis) > 0 {
				v.Children = append(v.Children, apis...)
			}
		}
		res = append(res, v)
	}
	return
}

func (s *sSysMenuService) Add(input sys_request.MenuAddOrEditReq) (err error) {
	var (
		sysMenuModel model.SysMenu
		parentIds    []int
		id           = input.Id
	)
	sysMenuModel, err = s.DealAddOrEdit(input)
	if err != nil {
		return
	}
	tx := easy_db.GetDb().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if id > 0 {
		err = tx.Model(model.SysMenu{}).Where("id =?", id).Save(&sysMenuModel).Error
	} else {
		err = tx.Model(model.SysMenu{}).Create(&sysMenuModel).Error
	}
	if err != nil {
		return
	}

	// 保存parent_ids
	parentIds, err = s.GetParentId(sysMenuModel.ParentId)
	parentIds = append(parentIds, sysMenuModel.Id)
	sysMenuModel.ParentIds = strings.Join(s.IntToStringArray(parentIds), ",")
	err = tx.Save(&sysMenuModel).Error
	if err != nil {
		return
	}

	// 保存apiRule
	err = s.saveApiRule(sysMenuModel.Id, input.ApisId)
	if err != nil {
		return
	}
	return
}

func (s *sSysMenuService) DealAddOrEdit(input sys_request.MenuAddOrEditReq) (sysMenuModel model.SysMenu, err error) {
	var (
		id         = input.Id
		parentId   = input.ParentId
		icon       = input.Icon
		title      = input.Title
		sort       = input.Sort
		path       = input.Path
		menuName   = input.MenuName
		permission = input.Permission
		menuType   = input.MenuType
		isFrame    = input.IsFrame
		visible    = input.Visible
		apisId     = input.ApisId
	)

	if id > 0 {
		err = easy_db.GetDb().First(&sysMenuModel, id).Error
		if err != nil {
			return
		}
		if len(sysMenuModel.ApisId) > 0 {
			tmpApisId := strings.Split(sysMenuModel.ApisId, ",")
			for _, val := range tmpApisId {
				apiId, _ := strconv.Atoi(val)
				apisId = append(apisId, apiId)
			}
		}

	}
	apisStr := s.IntToStringArray(apisId)

	switch menuType {
	case "M": // 目录
		if menuName == "" {
			err = errors.New("路由名称不能为空")
			return
		}

		if path == "" {
			err = errors.New("路径不能为空")
			return
		}
		sysMenuModel.MenuName = menuName
		sysMenuModel.Path = path
		sysMenuModel.IsFrame = isFrame
		sysMenuModel.Visible = visible
		sysMenuModel.Path = path
	case "C": // 菜单
		if menuName == "" {
			err = errors.New("路由名称不能为空")
			return
		}

		if path == "" {
			err = errors.New("路径不能为空")
			return
		}
		if permission == "" {
			err = errors.New("权限标识不能为空")
			return
		}
		sysMenuModel.MenuName = menuName
		sysMenuModel.Path = path
		sysMenuModel.IsFrame = isFrame
		sysMenuModel.Visible = visible
		sysMenuModel.Path = path
		sysMenuModel.Permission = permission
		sysMenuModel.ApisId = strings.Join(apisStr, ",")
	case "F": // 功能
		if permission == "" {
			err = errors.New("权限标识不能为空")
			return
		}
		sysMenuModel.Permission = permission
		sysMenuModel.ApisId = strings.Join(apisStr, ",")
	default:
		err = errors.New("格式异常")
		return
	}

	if icon != "" {
		sysMenuModel.Icon = icon
	}
	sysMenuModel.MenuType = menuType
	sysMenuModel.ParentId = parentId
	sysMenuModel.Title = title
	sysMenuModel.Sort = sort
	sysMenuModel.Operator = utils.GetUserName(s.ctx)
	return
}

func (s *sSysMenuService) Edit(input sys_request.MenuAddOrEditReq) (err error) {
	var (
		sysMenuModel model.SysMenu
	)

	sysMenuModel, err = s.DealAddOrEdit(input)
	if err != nil {
		return
	}
	err = easy_db.GetDb().Updates(&sysMenuModel).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysMenuService) GetParentId(parentId int) (parentIds []int, err error) {
	var tmp model.SysMenu
	err = easy_db.GetDb().Model(tmp).Where("id = ?", parentId).Scan(&tmp).Error
	if err != nil {
		return
	}
	if tmp.ParentId != 0 {
		var tmpParentId []int
		tmpParentId, err = s.GetParentId(tmp.ParentId)

		if err != nil {
			return
		}
		parentIds = append(parentIds, tmpParentId...)
	}
	parentIds = append(parentIds, tmp.Id)
	return
}

func (s *sSysMenuService) IntToStringArray(parentIds []int) (parentIdStr []string) {
	for _, val := range parentIds {
		res := strconv.Itoa(val)
		parentIdStr = append(parentIdStr, res)
	}
	return
}

func (s *sSysMenuService) saveApiRule(menuId int, apis []int) (err error) {
	var (
		sysApiList []model.SysApi
		apisStr    []string

		noApi               []string
		sysMenuApiRuleModel model.SysMenuApiRule
	)
	apisStr = s.IntToStringArray(apis)
	err = easy_db.GetDb().Model(model.SysApi{}).Where(fmt.Sprintf("id in (%v)", strings.Join(apisStr, ","))).Scan(&sysApiList).Error
	if err != nil {
		return
	}
	for _, item := range sysApiList {
		var flag bool
		for _, id := range apis {
			if item.ID == uint64(id) {
				flag = true
				break
			}
		}
		if !flag {
			noApi = append(noApi, string(item.ID))
			continue
		}
		sysMenuApiRuleModel.SysApiID = item.ID
		sysMenuApiRuleModel.SysMenuID = uint64(menuId)
		err = easy_db.GetDb().Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&sysMenuApiRuleModel).Error
		if err != nil {
			break
		}
	}

	if err != nil {
		return
	}
	if len(noApi) > 0 {
		err = errors.New(fmt.Sprintf("接口ID：%v，不能存在", strings.Join(noApi, ",")))
		return
	}

	return
}
