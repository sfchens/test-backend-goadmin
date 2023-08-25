package sys_service

import (
	"csf/app/admin/model/sys_model"
	"csf/app/admin/request/sys_req"
	"csf/core/mysql/model"
	"csf/library/easy_db"
	"csf/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type sSysDeptService struct {
	ctx *gin.Context
}

func NewSysDeptService(ctx *gin.Context) *sSysDeptService {
	return &sSysDeptService{
		ctx: ctx,
	}
}

func (s *sSysDeptService) AddOrEdit(input sys_req.DeptAddOrEditReq) (err error) {

	var (
		sysDeptModel model.SysDept
	)
	sysDeptModel, err = s.DealAddOrEdit(input)
	fmt.Printf("sysDeptModel:  %+v\n", sysDeptModel)
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

func (s *sSysDeptService) DealAddOrEdit(input sys_req.DeptAddOrEditReq) (sysDept model.SysDept, err error) {
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
	sysDept.Operator = utils.GetUserName(s.ctx)
	return
}

func (s *sSysDeptService) Edit(input sys_req.DeptAddOrEditReq) (err error) {
	var (
		id           = input.Id
		sysDeptModel model.SysDept
	)
	sysDeptModel, err = s.DealAddOrEdit(input)
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

func (s *sSysDeptService) Delete(input sys_req.DeptDeleteReq) (err error) {
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

	err = s.DeleteDeal(tran, idTmp)
	if err != nil {
		return
	}
	return
}

func (s *sSysDeptService) DeleteDeal(tran *gorm.DB, ids []string) (err error) {
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
		err = s.DeleteDeal(tran, pids)
		if err != nil {
			return
		}
	}
	return
}

func (s *sSysDeptService) GetQuery(input sys_req.DeptTreeListReq) *gorm.DB {
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
	return model
}

func (s *sSysDeptService) TreeList(input sys_req.DeptTreeListReq) (out sys_req.DeptTreeListRes, err error) {

	var (
		page     = input.Page
		pageSize = input.PageSize
		order    = input.Order

		sysMenuListTmp []sys_model.DeptTreeListItem
	)

	model := s.GetQuery(input)
	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model.Offset((page - 1) * pageSize).Preload("Children").Limit(pageSize).Order(order).Scan(&sysMenuListTmp).Error
	if err != nil {
		return
	}
	out.List = s.TreeListItem(sysMenuListTmp)
	return
}

func (s *sSysDeptService) TreeListItem(list []sys_model.DeptTreeListItem) (out []sys_model.DeptTreeListItem) {
	for _, v := range list {
		v.Label = v.Name
		model := easy_db.GetDb().Model(model.SysDept{}).Preload("Children").Where("parent_id = ?", v.ID)
		model.Order("sort desc").Scan(&v.Children)
		if len(v.Children) > 0 {
			v.Children = s.TreeListItem(v.Children)
		}
		out = append(out, v)
	}
	return
}

func (s *sSysDeptService) GetOne(input sys_req.DeptGetOneReq) (out sys_req.DeptGetOneRes, err error) {
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

func (s *sSysDeptService) DeleteMulti(input sys_req.DeptDeleteMultiReq) (err error) {

	var (
		ids = input.Ids

		errNew []string
	)
	for _, id := range ids {
		newInput := sys_req.DeptDeleteReq{
			Id: id,
		}
		err = s.Delete(newInput)
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
