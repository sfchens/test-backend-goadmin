package product

import (
	"csf/core/mysql/model"
	"csf/core/query/product_query"
	"csf/core/service"
	"csf/library/easy_db"
	"csf/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
)

func init() {
	service.RegisterNewProductCategory(NewCategoryService())
}

type sCategory struct{}

func NewCategoryService() *sCategory {
	return &sCategory{}
}
func (s *sCategory) AddOrEdit(ctx *gin.Context, input product_query.CategoryAddOrEditInput) (err error) {
	var (
		id     = input.Id
		pid    = input.Pid
		name   = input.Name
		sort   = input.Sort
		pic    = input.Pic
		bigPic = input.BigPic
		isShow = input.IsShow

		categoryModel model.ProductCategory
	)

	if id != 0 {
		err = easy_db.GetDb().Model(categoryModel).Where("id=?", id).Find(&categoryModel).Error
		if err != nil {
			return
		}
	}

	tx := easy_db.GetDb().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if pid != 0 {
		var existCount int64
		err = easy_db.GetDb().Model(categoryModel).Where("id=?", pid).Count(&existCount).Error
		if err != nil {
			return
		}
		if existCount <= 0 {
			err = errors.New("父级分类不存在")
			return
		}
	}

	categoryModel.PID = pid
	categoryModel.Name = name
	categoryModel.Sort = sort
	categoryModel.Pic = pic
	categoryModel.BigPic = bigPic
	categoryModel.IsShow = isShow
	categoryModel.Operator = utils.GetUserName(ctx)
	if categoryModel.ID <= 0 {
		categoryModel.CreatedAt = time.Now()
	}
	err = tx.Save(&categoryModel).Error
	if err != nil {
		return
	}
	categoryModel.PIDs = fmt.Sprintf("%v,%v", categoryModel.PID, categoryModel.ID)
	err = tx.Save(&categoryModel).Error
	if err != nil {
		return
	}

	return
}

func (s *sCategory) List(ctx *gin.Context, input product_query.CategoryListInput) (out product_query.CategoryListOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize

		sysMenuListTmp []product_query.CategoryTreeListItem
	)

	model := s.getQuery(input)
	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model.Offset((page - 1) * pageSize).Preload("Children").Limit(pageSize).Scan(&sysMenuListTmp).Error
	if err != nil {
		return
	}
	out.List = s.TreeListItem(ctx, sysMenuListTmp)
	return
}

func (s *sCategory) getQuery(input product_query.CategoryListInput) *gorm.DB {
	var (
		pids          = input.Pids
		isShow        = input.IsShow
		categoryModel model.ProductCategory
	)
	m := easy_db.GetDb().Model(categoryModel)

	if len(pids) > 0 {
		m.Where("id=?", pids[len(pids)-1])
	} else {
		m.Where("pid=?", 0)
	}

	if isShow != -1 {
		m.Where("is_show=?", isShow)
	}
	return m
}

func (s *sCategory) TreeListItem(ctx *gin.Context, list []product_query.CategoryTreeListItem) (out []product_query.CategoryTreeListItem) {
	for _, v := range list {
		v.Label = v.Name
		v.Value = v.Id
		model := easy_db.GetDb().Model(model.ProductCategory{}).Preload("Children").Where("pid = ?", v.Id)
		model.Order("sort desc").Scan(&v.Children)
		if len(v.Children) > 0 {
			v.Children = s.TreeListItem(ctx, v.Children)
		}
		out = append(out, v)
	}
	return
}

func (s *sCategory) DeleteBatch(ctx *gin.Context, input product_query.CategoryDeleteBatchInput) (err error) {
	var (
		ids          = input.Ids
		categoryList []model.ProductCategory

		m model.ProductCategory
	)

	idsStr := strings.Join(utils.IntToStringArray(ids), ",")

	err = easy_db.GetDb().Model(m).Where("id in (?)", idsStr).Scan(&categoryList).Error
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

	for _, item := range categoryList {
		err = s.Delete(tx, item.ID)
		if err != nil {
			break
		}
	}
	return
}

func (s *sCategory) Delete(tx *gorm.DB, id int) (err error) {
	var categoryInfo model.ProductCategory
	err = tx.Find(&categoryInfo, id).Error
	if err != nil {
		return
	}
	err = tx.Delete(&categoryInfo, id).Error
	if err != nil {
		return
	}
	var pidList []model.ProductCategory
	err = tx.Model(model.ProductCategory{}).Where("pid=?", id).Scan(&pidList).Error
	if err != nil {
		return
	}
	for _, item := range pidList {
		err = s.Delete(tx, item.ID)
		if err != nil {
			return
		}
	}
	return
}
