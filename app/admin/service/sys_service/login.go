package sys_service

import (
	"csf/app/admin/model/sys_model"
	"csf/app/admin/request/sys_request"
	"csf/common/mysql/model"
	"csf/library/custom_session"
	"csf/library/db"
	"csf/library/global"
	"csf/library/my_jwt"
	"csf/utils"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
)

type SysLogin struct {
	ctx *gin.Context
}

func NewSysLoginService(ctx *gin.Context) *SysLogin {
	return &SysLogin{ctx: ctx}
}

func (s *SysLogin) Login(input sys_model.LoginInput) (out sys_request.LoginRes, err error) {
	var (
		username = input.Username
		password = input.Password

		adminInfoModel model.SysAdmin
	)
	err = db.GetDb().Where("username=?", username).Find(&adminInfoModel).Error
	if err != nil {
		return
	}
	if !utils.BcryptCheck(password, adminInfoModel.Password) {
		err = errors.New("账号或者密码错误")
		return
	}

	if adminInfoModel.Status == 0 {
		err = errors.New("账号未启用")
		return
	}

	if adminInfoModel.Status == 2 {
		err = errors.New("账号已封禁")
		return
	}
	var tokenInfo sys_model.TokenInfoOut
	tokenInfo, err = s.CreateToken(adminInfoModel)

	var adminInfo sys_request.AdminInfo
	utils.StructToStruct(adminInfoModel, &adminInfo)
	out = sys_request.LoginRes{
		AdminInfo: adminInfo,
		TokenInfo: tokenInfo,
	}

	// 保存session
	err = custom_session.NewCustomSession(s.ctx).Set(global.LoginTypeKey, global.LoginTypeAdmin)
	if err != nil {
		return
	}
	jsonStr, _ := json.Marshal(adminInfoModel)
	err = custom_session.NewCustomSession(s.ctx).Set(global.UserInfoKey, string(jsonStr))
	if err != nil {
		return
	}
	return
}

func (s *SysLogin) CreateToken(input model.SysAdmin) (tokenInfoOUt sys_model.TokenInfoOut, err error) {
	var token string
	baseClaims := my_jwt.BaseClaims{
		Id:       int(input.ID),
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}
	myBaseClaims := my_jwt.NewJWT().CreateClaims(baseClaims)
	token, err = my_jwt.NewJWT().CreateToken(myBaseClaims)
	if err != nil {
		return
	}

	tokenInfoOUt = sys_model.TokenInfoOut{
		Token:     token,
		ExpiresAt: myBaseClaims.StandardClaims.ExpiresAt,
	}
	return
}

func (s *SysLogin) Logout() (err error) {
	err = custom_session.NewCustomSession(s.ctx).Delete(global.LoginTypeKey)
	if err != nil {
		return
	}
	err = custom_session.NewCustomSession(s.ctx).Delete(global.UserInfoKey)
	if err != nil {
		return
	}
	return
}