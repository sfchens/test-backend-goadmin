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
	"strings"
)

func init() {
	service.RegisterNewDept(NewSysDeptService())
}

type sSysDeptService struct{}

func NewSysDeptService() *sSysDeptService {
	return &sSysDeptService{}
}

func (s *sSysDeptService) AddOrEdit(ctx *gin.Context, input sys_query.DeptAddOrEditInput) (err error) {

	var (
		sysDeptModel model.SysDept
	)
	sysDeptModel, err = s.DealAddOrEdit(ctx, input)

	if err != nil {
		return
	}
	if sysDeptModel.ID > 0 {
		err = easy_db.GetDb().Updates(sysDeptModel).Error
	} else {
		err = easy_db.GetDb().Create(&sysDeptModel).Error

	}
	if err != nil {
		return
	}

	return
}

func (s *sSysDeptService) DealAddOrEdit(ctx *gin.Context, input sys_query.DeptAddOrEditInput) (sysDept model.SysDept, err error) {
	var (
		id       = input.Id
		parentId = input.ParentId
		name     = input.Name
		leader   = input.Leader
		sort     = input.Sort
		phone    = input.Phone
		email    = input.Email
		status   = input.Status
	)

	if parentId > 0 {
		var parentIdCount int64
		parentModel := easy_db.GetDb().Model(sysDept).Where("id = ?", parentId)
		err = parentModel.Count(&parentIdCount).Error
		if err != nil {
			return
		}

		if parentIdCount <= 0 {
			err = errors.New("上级部门不存在")
			return
		}
	}

	if id > 0 {
		err = easy_db.GetDb().First(&sysDept, id).Error
		if err != nil {
			return
		}
	}

	if phone != "" {
		sysDept.Phone = phone
	}
	if email != "" {
		sysDept.Email = email
	}
	sysDept.ParentId = parentId
	sysDept.Name = name
	sysDept.Leader = leader
	sysDept.Sort = sort
	sysDept.Status = status
	sysDept.Operator = utils.GetUserName(ctx)
	return
}

func (s *sSysDeptService) Edit(ctx *gin.Context, input sys_query.DeptAddOrEditInput) (err error) {
	var (
		id           = input.Id
		sysDeptModel model.SysDept
	)
	sysDeptModel, err = s.DealAddOrEdit(ctx, input)
	if err != nil {
		return
	}
	if id <= 0 {
		err = errors.New("参数异常")
		return
	}
	sysDeptModel.ID = id
	err = easy_db.GetDb().Updates(sysDeptModel).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysDeptService) Delete(ctx *gin.Context, input sys_query.DeptDeleteInput) (err error) {
	var (
		id           = input.Id
		sysDeptModel model.SysDept
	)

	err = easy_db.GetDb().First(&sysDeptModel, id).Error
	if err != nil {
		return
	}

	tran := easy_db.GetDb().Begin()
	defer func() {
		if err != nil {
			tran.Rollback()
		} else {
			tran.Commit()
		}

	}()

	idTmp := []string{fmt.Sprintf("%d", id)}

	err = s.DeleteDeal(ctx, tran, idTmp)
	if err != nil {
		return
	}
	return
}

func (s *sSysDeptService) DeleteDeal(ctx *gin.Context, tran *gorm.DB, ids []string) (err error) {
	var sysDeptModel model.SysDept
	// 删除
	models := tran.Where("id in (?)", strings.Join(ids, ","))
	err = models.Delete(&sysDeptModel, ids).Error
	if err != nil {
		return
	}

	// 父类
	var sysDeptParentList []model.SysDept
	err = tran.Model(sysDeptModel).Where("parent_id in (?)", strings.Join(ids, ",")).Scan(&sysDeptParentList).Error
	if err != nil {
		return
	}
	if len(sysDeptParentList) > 0 {
		var pids []string
		for _, item := range sysDeptParentList {
			pids = append(pids, fmt.Sprintf("%d", item.ID))
		}
		err = s.DeleteDeal(ctx, tran, pids)
		if err != nil {
			return
		}
	}
	return
}

func (s *sSysDeptService) getQuery(input sys_query.DeptTreeListInput) *gorm.DB {
	var (
		name     = input.Name
		status   = input.Status
		parentId = input.ParentId

		sysDeptModel model.SysDept
	)
	model := easy_db.GetDb().Model(sysDeptModel)
	if name != "" {
		model.Where(fmt.Sprintf("name like '%%%s%%'", name))
	}
	if status > 0 {
		model.Where("status = ?", status)
	}
	if parentId >= 0 {
		model.Where("parent_id = ?", parentId)
	}
	fmt.Printf("input: %+v\n", input)
	return model
}

func (s *sSysDeptService) TreeList(ctx *gin.Context, input sys_query.DeptTreeListInput) (out sys_query.DeptTreeListOut, err error) {

	var (
		page     = input.Page
		pageSize = input.PageSize
		order    = input.Order

		sysMenuListTmp []sys_query.DeptTreeListItem
	)

	model := s.getQuery(input)
	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model.Offset((page - 1) * pageSize).Preload("Children").Limit(pageSize).Order(order).Scan(&sysMenuListTmp).Error
	if err != nil {
		return
	}
	out.List = s.TreeListItem(ctx, sysMenuListTmp)
	return
}

func (s *sSysDeptService) TreeListItem(ctx *gin.Context, list []sys_query.DeptTreeListItem) (out []sys_query.DeptTreeListItem) {
	for _, v := range list {
		v.Label = v.Name
		model := easy_db.GetDb().Model(model.SysDept{}).Preload("Children").Where("parent_id = ?", v.ID)
		model.Order("sort desc").Scan(&v.Children)
		if len(v.Children) > 0 {
			v.Children = s.TreeListItem(ctx, v.Children)
		}
		out = append(out, v)
	}
	return
}

func (s *sSysDeptService) GetOne(ctx *gin.Context, input sys_query.DeptGetOneInput) (out sys_query.DeptGetOneOut, err error) {
	var (
		id           = input.Id
		sysDeptModel model.SysDept
	)

	err = easy_db.GetDb().Find(&sysDeptModel, id).Error
	if err != nil {
		return
	}

	out.SysDept = sysDeptModel
	return
}

func (s *sSysDeptService) DeleteMulti(ctx *gin.Context, input sys_query.DeptDeleteMultiInput) (err error) {

	var (
		ids = input.Ids

		errNew []string
	)
	for _, id := range ids {
		newInput := sys_query.DeptDeleteInput{
			Id: id,
		}
		err = s.Delete(ctx, newInput)
		if err != nil {
			errNew = append(errNew, fmt.Sprintf("序号： %v 删除失败, 错误信息： %v", id, err.Error()))
		}
	}

	if len(errNew) > 0 {
		err = errors.New(strings.Join(errNew, "\n,"))
		return
	}

	return
}
