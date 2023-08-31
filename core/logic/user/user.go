package user

import (
	"csf/core/mysql/model"
	"csf/core/query/user_query"
	"csf/core/service"
	"csf/library/easy_db"
	"csf/library/easy_validator"
	"csf/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sUser struct{}

func init() {
	service.RegisterNewUser(NewUserService())
}
func NewUserService() *sUser {
	return &sUser{}
}

func (s *sUser) AddOrEdit(ctx *gin.Context, input user_query.UserAddOrEditInput) (err error) {
	var (
		id       = input.Id
		username = input.Username
		realname = input.Realname
		email    = input.Email
		phone    = input.Phone
		password = input.Password

		userModel model.SysUser
	)
	err = easy_validator.NewValidator().Validate(input)
	if err != nil {
		return
	}

	if id > 0 {
		err = easy_db.GetDb().Model(userModel).Find(&userModel, id).Error
		if err != nil {
			return
		}
	}

	var count int64
	if err = easy_db.GetDb().Model(userModel).Where("username=?", username).Count(&count).Error; err != nil {
		return
	}

	if count > 0 {
		err = errors.New("账号已存在")
		return
	}

	userModel.Username = username
	userModel.Realname = realname
	userModel.Status = 1
	userModel.Operator = utils.GetUserName(ctx)
	if email != "" {
		if !utils.IsEmail(email) {
			err = errors.New("手机号码格式错误")
			return
		}
		if s.IsExitEmail(id, email) {
			err = errors.New("邮箱已存在")
			return
		}
		userModel.Email = email

	}
	if phone != "" {
		if !utils.IsPhone(phone) {
			err = errors.New("手机号码格式错误")
			return
		}
		if s.IsExitPhone(id, phone) {
			err = errors.New("邮箱已存在")
			return
		}
		userModel.Phone = phone
	}
	if password != "" {
		userModel.Password = utils.BcryptHash(password)
	}

	if userModel.ID > 0 {
		err = easy_db.GetDb().Save(&userModel).Error
	} else {
		err = easy_db.GetDb().Create(&userModel).Error
	}
	if err != nil {
		return
	}
	return
}

func (s *sUser) IsExitEmail(id int, email string) bool {

	exitEmailModel := easy_db.GetDb().Model(model.SysUser{}).Where("email=?", email)
	if id > 0 {
		exitEmailModel.Where("id != ?", id)
	}

	var exitCount int64
	exitEmailModel.Count(&exitCount)
	return exitCount > 0
}

func (s *sUser) IsExitPhone(id int, phone string) bool {

	exitEmailModel := easy_db.GetDb().Model(model.SysUser{}).Where("phone=?", phone)
	if id > 0 {
		exitEmailModel.Where("id != ?", id)
	}

	var exitCount int64
	exitEmailModel.Count(&exitCount)
	return exitCount > 0
}

func (s *sUser) ResetPwd(ctx *gin.Context, input user_query.UserResetPwdInput) (err error) {
	var (
		id       = input.Id
		password = input.Password

		userModel model.SysUser
	)
	err = easy_validator.NewValidator().Validate(input)
	if err != nil {
		return
	}

	err = easy_db.GetDb().Model(userModel).Find(&userModel, id).Error
	if err != nil {
		return
	}

	userModel.Password = utils.BcryptHash(password)
	userModel.Operator = utils.GetUserName(ctx)
	err = easy_db.GetDb().Save(&userModel).Error
	if err != nil {
		return
	}

	return
}

func (s *sUser) SetStatus(ctx *gin.Context, input user_query.UserSetStatusInput) (err error) {
	var (
		id     = input.Id
		status = input.Status

		userModel model.SysUser
	)

	err = easy_validator.NewValidator().Validate(input)
	if err != nil {
		return
	}

	err = easy_db.GetDb().Model(userModel).Find(&userModel, id).Error
	if err != nil {
		return
	}

	if userModel.Status == status {
		err = errors.New("状态异常")
		return
	}
	userModel.Operator = utils.GetUserName(ctx)
	err = easy_db.GetDb().Save(&userModel).Error
	if err != nil {
		return
	}

	return
}

func (s *sUser) List(ctx *gin.Context, input user_query.UserListInput) (total int64, out []user_query.UserListItem, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
	)

	m := s.getQuery(input)
	err = m.Count(&total).Error
	if err != nil {
		return
	}

	var userList []model.SysUser
	err = m.Offset((page - 1) * pageSize).Limit(pageSize).Scan(&userList).Error
	if err != nil {
		return
	}

	for _, item := range userList {
		var tmp user_query.UserListItem
		utils.StructToStruct(item, &tmp)

		out = append(out, tmp)
	}
	return
}

func (s *sUser) getQuery(input user_query.UserListInput) *gorm.DB {

	var (
		username = input.Username
		realname = input.Realname
		email    = input.Email
		phone    = input.Phone
		status   = input.Status

		userModel model.SysUser
	)

	model := easy_db.GetDb().Model(userModel)
	if username != "" {
		model.Where(fmt.Sprintf("username %%%v%%v", username))
	}
	if realname != "" {
		model.Where(fmt.Sprintf("realname %%%v%%v", realname))
	}
	if email != "" {
		model.Where("email = ?", email)
	}
	if phone != "" {
		model.Where("phone = ?", phone)
	}

	if status != -1 {
		model.Where("status = ?", status)
	}
	return model
}

func (s *sUser) GetInfo(ctx *gin.Context, id int) (out user_query.UserListItem, err error) {
	var (
		userModel model.SysUser
	)
	err = easy_db.GetDb().Model(userModel).Where("id=?", id).Scan(&userModel).Error
	if err != nil {
		return
	}

	if userModel.ID <= 0 {
		err = errors.New("数据不存在")
		return
	}

	utils.StructToStruct(userModel, &out)
	return
}
