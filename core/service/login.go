package service

import (
	"csf/core/mysql/model"
	"csf/core/query/login_query"
	"github.com/gin-gonic/gin"
)

var localLoginService loginServiceGroup

func NewLoginServiceGroup() loginServiceGroup {
	return localLoginService
}

type loginServiceGroup struct {
	AdminLoginService iAdminLogin
	H5LoginService    iH5Login
}

type (
	iAdminLogin interface {
		Login(ctx *gin.Context, input login_query.AdminLoginInput) (out login_query.AdminLoginOut, err error)
		CreateToken(ctx *gin.Context, input model.SysAdmin) (tokenInfoOUt login_query.TokenInfoOut, err error)
		Logout(ctx *gin.Context) (err error)
	}

	iH5Login interface {
		Login(ctx *gin.Context, input login_query.H5LoginInput) (out login_query.H5LoginOut, err error)
		Logout(ctx *gin.Context) (err error)
	}
)

func RegisterAdminLogin(i iAdminLogin) {
	localLoginService.AdminLoginService = i
}

func RegisterH5Login(i iH5Login) {
	localLoginService.H5LoginService = i
}
