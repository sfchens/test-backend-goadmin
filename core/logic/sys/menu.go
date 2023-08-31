package sys

import (
	"csf/core/mysql/model"
	"csf/core/query/sys_query"
	"csf/core/service"
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

func init() {
	service.RegisterNewMenu(NewSysMenuService())
}

type sSysMenuService struct {
}

func NewSysMenuService() *sSysMenuService {
	return &sSysMenuService{}
}

func (s *sSysMenuService) TreeList(ctx *gin.Context, input sys_query.MenuTreeListInput) (out sys_query.MenuTreeListOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize

		sysMenuListTmp []sys_query.MenuListItem

		inputTmp sys_query.MenuListInput
	)

	utils.StructToStruct(input, &inputTmp)
	inputTmp.MenuType = "M"
	model := s.getQuery(inputTmp)

	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model.Offset((page - 1) * pageSize).Preload("Children").Limit(pageSize).Scan(&sysMenuListTmp).Error
	if err != nil {
		return
	}
	list := s.GetMenuItem(ctx, sysMenuListTmp, false)
	out.List = list
	return
}

func (s *sSysMenuService) List(ctx *gin.Context, input sys_query.MenuListInput) (out sys_query.MenuListOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize

		sysMenuListTmp []sys_query.MenuListItem
	)

	model := s.getQuery(input)

	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model.Offset((page - 1) * pageSize).Preload("Children").Limit(pageSize).Order("sort ASC").Scan(&sysMenuListTmp).Error
	if err != nil {
		return
	}
	list := s.GetMenuItem(ctx, sysMenuListTmp, true)

	out.List = s.DealTreeList(ctx, list)
	return
}

func (s *sSysMenuService) DealTreeList(ctx *gin.Context, data []sys_query.MenuListItem) (res []sys_query.SysMenuListItem) {
	for _, item := range data {
		var tmpV sys_query.SysMenuListItem
		utils.StructToStruct(item.SysMenu, &tmpV)
		tmpV.Label = item.Title
		// 取出apis
		if item.ApisID != "" {
			var apiIds []int
			for _, apiId := range strings.Split(item.ApisID, ",") {
				tmp, _ := strconv.Atoi(apiId)
				apiIds = append(apiIds, tmp)
			}
			tmpV.ApisId = apiIds
			var apiList []model.SysAPI
			easy_db.GetDb().Model(model.SysAPI{}).Where(fmt.Sprintf("id in (%v)", item.ApisID)).Scan(&apiList)
			tmpV.SysApis = apiList
		} else {
			out, _ := s.GetApisByMenuId(ctx, []int{int(item.ID)})
			tmpV.ApisId = out.ApisId
			tmpV.SysApis = out.ApisList
		}
		if len(item.Children) > 0 {
			tmpV.Children = s.DealTreeList(ctx, item.Children)
		}
		if len(tmpV.ApisId) <= 0 {
			tmpV.ApisId = []int{}
		}
		res = append(res, tmpV)
	}
	return
}

func (s *sSysMenuService) GetApisByMenuId(ctx *gin.Context, ids []int) (out sys_query.GetApisByMenuIdOut, err error) {
	idsArr := utils.IntToStringArray(ids)
	idsStr := strings.Join(idsArr, ",")

	err = easy_db.GetDb().Model(model.SysAPI{}).Where(fmt.Sprintf("id in (%v)", idsStr)).Scan(&out.ApisList).Error
	if err != nil {
		return
	}

	for _, item := range out.ApisList {
		out.ApisId = append(out.ApisId, int(item.ID))
	}
	return
}

func (s *sSysMenuService) getQuery(input sys_query.MenuListInput) *gorm.DB {
	var (
		title    = input.Key
		visible  = input.Visible
		menuType = input.MenuType

		sysMenu model.SysMenu
	)

	model := easy_db.GetDb().Model(sysMenu)

	if title != "" {
		model.Where("title like '%?%'", title)
	}

	if visible != -1 {
		model.Where("visible = ?", visible)
	}

	if menuType != "" {
		model.Where("menu_type = ?", menuType)
	}
	model.Where("parent_id=?", 0)
	return model
}

func (s *sSysMenuService) GetMenuItem(ctx *gin.Context, list []sys_query.MenuListItem, isAll bool) (res []sys_query.MenuListItem) {
	for _, v := range list {
		model1 := easy_db.GetDb().Model(model.SysMenu{}).Preload("Children").Where("menu_type != ?", "F").Where("parent_id = ?", v.ID)
		model1.Order("sort ASC").Scan(&v.Children)
		if len(v.Children) > 0 {
			v.Children = s.GetMenuItem(ctx, v.Children, isAll)
		}
		if isAll {
			var apis []sys_query.MenuListItem
			easy_db.GetDb().Model(model.SysMenu{}).Preload("Children").
				Where("menu_type = ?", "F").
				Where("parent_id = ?", v.ID).Order("sort ASC").Scan(&apis)
			if len(apis) > 0 {
				v.Children = append(v.Children, apis...)
			}
		}
		res = append(res, v)
	}
	return
}

func (s *sSysMenuService) Add(ctx *gin.Context, input sys_query.MenuAddOrEditInput) (err error) {
	var (
		sysMenuModel model.SysMenu
		parentIds    []int
		id           = input.Id
	)
	sysMenuModel, err = s.DealAddOrEdit(ctx, input)
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
	parentIds, err = s.GetParentId(int(sysMenuModel.ParentID))
	parentIds = append(parentIds, int(sysMenuModel.ID))
	sysMenuModel.ParentIDs = strings.Join(utils.IntToStringArray(parentIds), ",")
	err = tx.Save(&sysMenuModel).Error
	if err != nil {
		return
	}
	if len(input.ApisId) > 0 {
		// 保存apiRule
		err = s.saveApiRule(int(sysMenuModel.ID), input.ApisId)
		if err != nil {
			return
		}
	}
	return
}

func (s *sSysMenuService) DealAddOrEdit(ctx *gin.Context, input sys_query.MenuAddOrEditInput) (sysMenuModel model.SysMenu, err error) {
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
		component  = input.Component
	)

	if id > 0 {
		err = easy_db.GetDb().First(&sysMenuModel, id).Error
		if err != nil {
			return
		}
		if len(sysMenuModel.ApisID) > 0 {
			tmpApisId := strings.Split(sysMenuModel.ApisID, ",")
			for _, val := range tmpApisId {
				apiId, _ := strconv.Atoi(val)
				apisId = append(apisId, apiId)
			}
		}

	}
	apisStr := utils.IntToStringArray(apisId)

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
		sysMenuModel.Component = component
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
		sysMenuModel.Component = component
		sysMenuModel.ApisID = strings.Join(apisStr, ",")
	case "F": // 功能
		if permission == "" {
			err = errors.New("权限标识不能为空")
			return
		}
		sysMenuModel.Permission = permission
		sysMenuModel.ApisID = strings.Join(apisStr, ",")
	default:
		err = errors.New("格式异常")
		return
	}

	if icon != "" {
		sysMenuModel.Icon = icon
	}
	sysMenuModel.MenuType = menuType
	sysMenuModel.ParentID = int64(parentId)
	sysMenuModel.Title = title
	sysMenuModel.Sort = sort
	sysMenuModel.Operator = utils.GetUserName(ctx)
	return
}

func (s *sSysMenuService) Edit(ctx *gin.Context, input sys_query.MenuAddOrEditInput) (err error) {
	var (
		sysMenuModel model.SysMenu
	)

	sysMenuModel, err = s.DealAddOrEdit(ctx, input)
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
	if tmp.ParentID != 0 {
		var tmpParentId []int
		tmpParentId, err = s.GetParentId(int(tmp.ParentID))

		if err != nil {
			return
		}
		parentIds = append(parentIds, tmpParentId...)
	}
	parentIds = append(parentIds, int(tmp.ID))
	return
}

func (s *sSysMenuService) saveApiRule(menuId int, apis []int) (err error) {
	var (
		sysApiList []model.SysAPI
		apisStr    []string

		noApi               []string
		sysMenuApiRuleModel model.SysMenuAPIRule
	)

	apisStr = utils.IntToStringArray(apis)
	fmt.Printf("apisStr:  %+v\n", apis)
	err = easy_db.GetDb().Model(model.SysAPI{}).Where(fmt.Sprintf("id in (%v)", strings.Join(apisStr, ","))).Scan(&sysApiList).Error
	if err != nil {
		return
	}
	for _, item := range sysApiList {
		var flag bool
		for _, id := range apis {
			if item.ID == int64(id) {
				flag = true
				break
			}
		}
		if !flag {
			noApi = append(noApi, string(item.ID))
			continue
		}
		sysMenuApiRuleModel.SysAPIID = item.ID
		sysMenuApiRuleModel.SysMenuID = int64(menuId)
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

func (s *sSysMenuService) GetMenuByRoleId(roleIds []int) (err error) {
	var (
		roleMenuList   []model.SysRoleMenu
		roleList       []model.SysRole
		menuApiList    []model.SysMenu
		permissionList []string
	)

	if len(roleIds) <= 0 {
		err = errors.New("获取菜单路由权限参数异常")
		return
	}

	roleIdStr := utils.IntToStringArray(roleIds)
	// 角色信息
	err = easy_db.GetDb().Model(model.SysRole{}).Where(fmt.Sprintf("id in (%v)", strings.Join(roleIdStr, ","))).Scan(&roleList).Error
	if err != nil {
		return
	}

	// 获取角色与Api信息
	err = easy_db.GetDb().Model(model.SysRoleMenu{}).Where(fmt.Sprintf("role_id in (%v)", strings.Join(roleIdStr, ","))).
		Scan(&roleMenuList).Error
	if err != nil {
		return
	}

	// 获取菜单信息
	var menuIds []string
	for _, item := range roleMenuList {
		if item.RoleID == 1 {
			permissionList = append(permissionList, "*:*:*")
		}
		menuIds = append(menuIds, fmt.Sprintf("%v", item.MenuID))
	}

	err = easy_db.GetDb().Model(model.SysMenu{}).Where(fmt.Sprintf("id in (%v)", strings.Join(menuIds, ","))).
		Scan(&menuApiList).Error
	if err != nil {
		return
	}
	// 获取permission
	for _, item := range menuApiList {
		permissionList = append(permissionList, item.Permission)
	}

	return
}
