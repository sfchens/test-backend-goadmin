package service

import (
	"csf/core/mysql/model"
	"csf/core/query/sys_query"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var localSysService = sysServiceGroup{}

func NewSysServiceGroup() sysServiceGroup {
	return localSysService
}

type sysServiceGroup struct {
	AdminService iAdmin
	ApiService   iApi
	DeptService  iDept
	MenuService  iMenu
	RoleService  iRole
}

type (
	iAdmin interface {
		Add(ctx *gin.Context, input sys_query.AdminAddOrEditInput) (err error)
		DealAddOrEdit(ctx *gin.Context, input sys_query.AdminAddOrEditInput) (sysAdminModel model.SysAdmin, err error)
		SetStatus(ctx *gin.Context, input sys_query.AdminSetStatusInput) (err error)
		List(ctx *gin.Context, input sys_query.AdminListInput) (out sys_query.AdminListOut, err error)
		GetAdminInfo(ctx *gin.Context) (adminModel model.SysAdmin, err error)
		ResetPwd(ctx *gin.Context, input sys_query.AdminResetPwdInput) (err error)
		SetRole(ctx *gin.Context, input sys_query.AdminSetRoleInput) (err error)
		DeleteBatch(ctx *gin.Context, input sys_query.AdminDeleteBatchInput) (err error)
	}

	iApi interface {
		List(ctx *gin.Context, input sys_query.ApiListInput) (out sys_query.ApiListOut, err error)
		Refresh() (err error)
		AddOrEdit(ctx *gin.Context, input sys_query.ApiEditInput) (err error)
		GetTag(ctx *gin.Context, input sys_query.ApiGetTagInput) (out sys_query.ApiGetTagOut, err error)
		DeleteMulti(ctx *gin.Context, input sys_query.ApiDeleteMultiInput) (err error)
	}

	iDept interface {
		AddOrEdit(ctx *gin.Context, input sys_query.DeptAddOrEditInput) (err error)
		DealAddOrEdit(ctx *gin.Context, input sys_query.DeptAddOrEditInput) (sysDept model.SysDept, err error)
		Edit(ctx *gin.Context, input sys_query.DeptAddOrEditInput) (err error)
		Delete(ctx *gin.Context, input sys_query.DeptDeleteInput) (err error)
		DeleteDeal(ctx *gin.Context, tran *gorm.DB, ids []string) (err error)
		TreeList(ctx *gin.Context, input sys_query.DeptTreeListInput) (out sys_query.DeptTreeListOut, err error)
		TreeListItem(ctx *gin.Context, list []sys_query.DeptTreeListItem) (out []sys_query.DeptTreeListItem)
		GetOne(ctx *gin.Context, input sys_query.DeptGetOneInput) (out sys_query.DeptGetOneOut, err error)
		DeleteMulti(ctx *gin.Context, input sys_query.DeptDeleteMultiInput) (err error)
	}

	iMenu interface {
		TreeList(ctx *gin.Context, input sys_query.MenuTreeListInput) (out sys_query.MenuTreeListOut, err error)
		List(ctx *gin.Context, input sys_query.MenuListInput) (out sys_query.MenuListOut, err error)
		DealTreeList(ctx *gin.Context, data []sys_query.MenuListItem) (res []sys_query.SysMenuListItem)
		GetApisByMenuId(ctx *gin.Context, ids []int) (out sys_query.GetApisByMenuIdOut, err error)
		GetMenuItem(ctx *gin.Context, list []sys_query.MenuListItem, isAll bool) (res []sys_query.MenuListItem)
		Add(ctx *gin.Context, input sys_query.MenuAddOrEditInput) (err error)
		DealAddOrEdit(ctx *gin.Context, input sys_query.MenuAddOrEditInput) (sysMenuModel model.SysMenu, err error)
		Edit(ctx *gin.Context, input sys_query.MenuAddOrEditInput) (err error)
		GetParentId(parentId int) (parentIds []int, err error)
		GetMenuByRoleId(roleIds []int) (err error)
	}

	iRole interface {
		AddOrEdit(ctx *gin.Context, input sys_query.RoleAddOrEditInput) (err error)
		SaveRoleMenu(ctx *gin.Context, tx *gorm.DB, roleId int, menuIds []int) (err error)
		List(ctx *gin.Context, input sys_query.RoleListInput) (out sys_query.RoleListOut, err error)
		DeleteBatch(ctx *gin.Context, input sys_query.RoleDeleteBatchInput) (err error)
		Delete(ctx *gin.Context, input sys_query.RoleDeleteInput) (err error)
	}
)

func RegisterNewAdmin(i iAdmin) {
	localSysService.AdminService = i
}

func RegisterNewApi(i iApi) {
	localSysService.ApiService = i
}

func RegisterNewDept(i iDept) {
	localSysService.DeptService = i
}

func RegisterNewMenu(i iMenu) {
	localSysService.MenuService = i
}

func RegisterNewRole(i iRole) {
	localSysService.RoleService = i
}
