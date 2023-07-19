package sys_service

import (
	"csf/app/admin/model/sys_model"
	"csf/app/admin/request/sys_request"
	"csf/common/mysql/model"
	"csf/library/db"
	"csf/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	out.List = s.GetMenuItem(sysMenuListTmp)
	return
}

func (s *sSysMenuService) GetQuery(input sys_request.MenuListReq) *gorm.DB {
	var (
		title  = input.Title
		isShow = input.IsShow

		sysMenu model.SysMenu
	)

	model := db.GetDb().Model(sysMenu)

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

func (s *sSysMenuService) GetMenuItem(list []sys_model.MenuListItem) (res []sys_model.MenuListItem) {
	for _, v := range list {
		model := db.GetDb().Model(model.SysMenu{}).Preload("Children").Where("menu_type!=?", "F").Where("parent_id = ?", v.Id)
		model.Order("sort desc").Scan(&v.Children)
		if len(v.Children) > 0 {
			v.Children = s.GetMenuItem(v.Children)
		}
		res = append(res, v)
	}
	return
}

func (s *sSysMenuService) Add(input sys_request.MenuAddOrEditReq) (err error) {
	var (
		sysMenuModel model.SysMenu
	)
	sysMenuModel, err = s.DealAddOrEdit(input)
	if err != nil {
		return
	}
	err = db.GetDb().Model(model.SysMenu{}).Create(&sysMenuModel).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysMenuService) DealAddOrEdit(input sys_request.MenuAddOrEditReq) (sysMenuModel model.SysMenu, err error) {
	var (
		id = input.Id
		//parentIds  = input.ParentIds
		icon  = input.Icon
		title = input.Title
		sort  = input.Sort
		//isShow     = input.IsShow
		path       = input.Path
		uniqueName = input.UniqueName
		uniqueAuth = input.UniqueAuth
		menuType   = input.MenuType
		isFrame    = input.IsFrame
	)

	if id > 0 {
		err = db.GetDb().First(&sysMenuModel, id).Error
		if err != nil {
			return
		}
	}

	//if len(parentIds) > 0 {
	//	sysMenuModel.ParentIds = strings.Join(parentIds, ",")
	//	parentId, _ := strconv.Atoi(parentIds[len(parentIds)-1])
	//	//sysMenuModel.ParentId = parentId
	//}
	if icon != "" {
		sysMenuModel.Icon = icon
	}
	switch menuType {
	case 1: // 目录
		if uniqueName == "" {
			err = errors.New("菜单名称必填")
			return
		}
		if isFrame <= 0 {
			err = errors.New("框架类型必选")
			return
		}
		//sysMenuModel.UniqueName = uniqueName
		//sysMenuModel.IsFrame = int8(isFrame)
		//sysMenuModel.IsShow = uint8(isShow)

	case 2: // 菜单
		if uniqueName == "" {
			err = errors.New("菜单名称必填")
			return
		}
		if uniqueAuth == "" {
			err = errors.New("权限唯一标识必填")
			return
		}
		//if apiUrl == "" {
		//	err = errors.New("接口地址必填")
		//	return
		//}
		//sysMenuModel.UniqueName = uniqueName
		//sysMenuModel.IsFrame = int8(isFrame)
		//sysMenuModel.IsShow = uint8(isShow)
		//sysMenuModel.UniqueAuth = uniqueAuth
		//sysMenuModel.ApiUrl = apiUrl

	case 3: // 按钮
		if uniqueAuth == "" {
			err = errors.New("权限唯一标识必填")
			return
		}
		//if apiUrl == "" {
		//	err = errors.New("接口地址必填")
		//	return
		//}
		//sysMenuModel.UniqueAuth = uniqueAuth
		//sysMenuModel.ApiUrl = apiUrl
	default:
		err = errors.New("类型错误")
		return
	}
	sysMenuModel.Title = title
	sysMenuModel.Path = path
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
	err = db.GetDb().Updates(&sysMenuModel).Error
	if err != nil {
		return
	}
	return
}
