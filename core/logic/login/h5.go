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
	service.RegisterH5Login(NewsH5LoginService())
}

type sH5Login struct{}

func NewsH5LoginService() *sH5Login {
	return &sH5Login{}
}

func (s *sH5Login) Login(ctx *gin.Context, input login_query.H5LoginInput) (out login_query.H5LoginOut, err error) {
	var (
		username = input.Username
		password = input.Password

		userInfoModel model.SysUser
	)
	err = easy_db.GetDb().Where("username=?", username).Find(&userInfoModel).Error
	if err != nil {
		return
	}
	if !utils.BcryptCheck(password, userInfoModel.Password) {
		err = errors.New("账号或者密码错误")
		return
	}

	if userInfoModel.Status == 0 {
		err = errors.New("账号未启用")
		return
	}

	if userInfoModel.Status == 2 {
		err = errors.New("账号已封禁")
		return
	}
	var tokenInfo login_query.TokenInfoOut
	tokenInfo, err = s.CreateToken(ctx, userInfoModel)

	var h5UserInfo login_query.H5UserInfo
	utils.StructToStruct(userInfoModel, &h5UserInfo)
	out = login_query.H5LoginOut{
		H5UserInfo: h5UserInfo,
		TokenInfo:  tokenInfo,
	}

	return
}

func (s *sH5Login) CreateToken(ctx *gin.Context, input model.SysUser) (tokenInfoOUt login_query.TokenInfoOut, err error) {
	var token string
	baseClaims := easy_auth.BaseClaims{
		Id:        int(input.ID),
		Username:  input.Username,
		Realname:  input.Realname,
		Email:     input.Email,
		Password:  input.Password,
		LoginType: global.ModuleH5,
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

func (s *sH5Login) Logout(ctx *gin.Context) (err error) {
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
