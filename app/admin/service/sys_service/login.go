package sys_service

import (
	"csf/app/admin/model/sys_model"
	"csf/app/admin/request/sys_req"
	"csf/common/mysql/model"
	"csf/library/easy_auth"
	"csf/library/easy_db"
	"csf/library/easy_session"
	"csf/library/global"
	"csf/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

type SysLogin struct {
	ctx *gin.Context
}

func NewSysLoginService(ctx *gin.Context) *SysLogin {
	return &SysLogin{ctx: ctx}
}

func (s *SysLogin) Login(input sys_model.LoginInput) (out sys_req.LoginRes, err error) {
	var (
		username = input.Username
		password = input.Password

		adminInfoModel model.SysAdmin
	)
	err = easy_db.GetDb().Where("username=?", username).Find(&adminInfoModel).Error
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

	var adminInfo sys_req.AdminInfo
	utils.StructToStruct(adminInfoModel, &adminInfo)
	out = sys_req.LoginRes{
		AdminInfo: adminInfo,
		TokenInfo: tokenInfo,
	}

	return
}

func (s *SysLogin) CreateToken(input model.SysAdmin) (tokenInfoOUt sys_model.TokenInfoOut, err error) {
	var token string
	baseClaims := easy_auth.BaseClaims{
		Id:        int(input.ID),
		Username:  input.Username,
		Realname:  input.Realname,
		Email:     input.Email,
		Password:  input.Password,
		LoginType: global.LoginTypeAdmin,
	}
	myBaseClaims := easy_auth.NewJWT().CreateClaims(baseClaims)
	token, err = easy_auth.NewJWT().CreateToken(myBaseClaims)
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
	err = easy_session.NewCustomSession(s.ctx).Delete(global.LoginTypeKey)
	if err != nil {
		return
	}
	err = easy_session.NewCustomSession(s.ctx).Delete(global.UserInfoKey)
	if err != nil {
		return
	}
	return
}
