package sys_service

import (
	"csf/app/admin/request/sys_request"
	"csf/common/mysql/model"
	"csf/library/db"
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

func (s *sSysDeptService) Add(input sys_request.DeptAddOrEditReq) (err error) {

	var (
		sysDeptModel model.SysDept
	)
	sysDeptModel, err = s.DealAddOrEdit(input)
	fmt.Printf("sysDeptModel:  %+v\n", sysDeptModel)
	if err != nil {
		return
	}
	err = db.GetDb().Create(&sysDeptModel).Error
	if err != nil {
		return
	}

	return
}

func (s *sSysDeptService) DealAddOrEdit(input sys_request.DeptAddOrEditReq) (sysDept model.SysDept, err error) {
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

	var parentIdCount int64
	parentModel := db.GetDb().Model(sysDept).Where("id = ?", parentId)
	err = parentModel.Count(&parentIdCount).Error
	if err != nil {
		return
	}

	if parentIdCount <= 0 {
		err = errors.New("上级部门不存在")
		return
	}

	if id > 0 {
		err = db.GetDb().First(&sysDept, id).Error
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

func (s *sSysDeptService) Edit(input sys_request.DeptAddOrEditReq) (err error) {
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
	err = db.GetDb().Updates(sysDeptModel).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysDeptService) Delete(input sys_request.DeptDeleteReq) (err error) {
	var (
		id           = input.Id
		sysDeptModel model.SysDept
	)

	err = db.GetDb().First(&sysDeptModel, id).Error
	if err != nil {
		return
	}

	tran := db.GetDb().Begin()
	defer func() {
		if err != nil {
			tran.Rollback()
		}
		tran.Commit()
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

func (s *sSysDeptService) List(input sys_request.DeptListReq) (out sys_request.DeptListRes, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
	)
	model := s.GetQuery(input)
	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}
	err = model.Offset((page - 1) * pageSize).Limit(pageSize).Scan(&out.List).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysDeptService) GetQuery(input sys_request.DeptListReq) *gorm.DB {
	var (
		name   = input.Name
		status = input.Status

		sysDeptModel model.SysDept
	)
	model := db.GetDb().Model(sysDeptModel)
	if name != "" {
		model.Where(fmt.Sprintf("name like '%%%s%%'", name))
	}
	fmt.Printf("status: %v\n", status)
	if status != -1 {
		model.Where("status = ?", status)
	}
	return model
}
