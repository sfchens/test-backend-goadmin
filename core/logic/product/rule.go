package product

import (
	"csf/core/mysql/model"
	"csf/core/query/product_query"
	"csf/core/service"
	"csf/library/easy_db"
	"csf/library/easy_validator"
	"csf/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func init() {
	service.RegisterNewProductRule(NewSysAdminService())
}

type sRule struct{}

func NewSysAdminService() *sRule {
	return &sRule{}
}

func (s *sRule) Add(ctx *gin.Context, input product_query.RuleAddOrEditInput) (err error) {
	var (
		id    = input.Id
		name  = input.Name
		value = input.Value

		ruleModel model.ProductRule
	)

	err = easy_validator.NewValidator().Validate(input)
	if err != nil {
		return
	}

	if id > 0 {
		err = easy_db.GetDb().Model(ruleModel).Find(&ruleModel, id).Error
		if err != nil {
			return
		}
	}

	ruleModel.Name = name
	ruleModel.Value = utils.ToJson(value)
	ruleModel.Operator = utils.GetUserName(ctx)
	if ruleModel.ID <= 0 {
		ruleModel.CreatedAt = time.Now()
		err = easy_db.GetDb().Model(ruleModel).Create(&ruleModel).Error
	} else {
		err = easy_db.GetDb().Model(ruleModel).Save(&ruleModel).Error
	}
	if err != nil {
		return
	}
	return
}

func (s *sRule) List(ctx *gin.Context, input product_query.RuleListInput) (out product_query.RuleListOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
		name     = input.Name

		ruleModel   model.ProductRule
		ruleListTmp []model.ProductRule
	)
	fmt.Printf("input: %+v\n", input)
	m := easy_db.GetDb().Model(&ruleModel)
	if name != "" {
		m.Where(fmt.Sprintf("name like '%%%v%%'", name))
	}

	err = m.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = m.Offset((page - 1) * pageSize).Limit(pageSize).Scan(&ruleListTmp).Error
	if err != nil {
		return
	}

	for _, item := range ruleListTmp {
		var tmp product_query.RuleModelOut
		utils.StructToStruct(item, &tmp)
		_ = json.Unmarshal([]byte(item.Value), &tmp.Value)
		out.List = append(out.List, tmp)
	}
	return
}
