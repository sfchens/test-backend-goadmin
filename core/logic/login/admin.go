package login

import (
	"csf/core/mysql/model"
	"csf/core/query/login_query"
	"csf/core/service"
	"csf/library/easy_auth"
	"csf/library/easy_db"
	"csf/library/easy_session"
	"csf/library/global"
	"csf/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func init() {
	service.RegisterAdminLogin(NewsAdminLoginService())
}

type sAdminLogin struct{}

func NewsAdminLoginService() *sAdminLogin {
	return &sAdminLogin{}
}

func (s *sAdminLogin) Login(ctx *gin.Context, input login_query.AdminLoginInput) (out login_query.AdminLoginOut, err error) {
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
	var tokenInfo login_query.TokenInfoOut
	tokenInfo, err = s.CreateToken(ctx, adminInfoModel)

	var adminInfo login_query.AdminInfo
	utils.StructToStruct(adminInfoModel, &adminInfo)
	out = login_query.AdminLoginOut{
		AdminInfo: adminInfo,
		TokenInfo: tokenInfo,
	}

	return
}

func (s *sAdminLogin) CreateToken(ctx *gin.Context, input model.SysAdmin) (tokenInfoOUt login_query.TokenInfoOut, err error) {
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

	tokenInfoOUt = login_query.TokenInfoOut{
		Token:     token,
		ExpiresAt: myBaseClaims.StandardClaims.ExpiresAt,
	}
	return
}

func (s *sAdminLogin) Logout(ctx *gin.Context) (err error) {
	err = easy_session.NewCustomSession(ctx).Delete(global.LoginTypeKey)
	if err != nil {
		return
	}
	err = easy_session.NewCustomSession(ctx).Delete(global.UserInfoKey)
	if err != nil {
		return
	}
	return
}
