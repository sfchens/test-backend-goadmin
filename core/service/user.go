package service

import (
	"csf/core/query/user_query"
	"github.com/gin-gonic/gin"
)

var localUserService userServiceGroup

func NewUserServiceGroup() userServiceGroup {
	return localUserService
}

type userServiceGroup struct {
	UserService iUser
}

type (
	iUser interface {
		AddOrEdit(ctx *gin.Context, input user_query.UserAddOrEditInput) (err error)
		ResetPwd(ctx *gin.Context, input user_query.UserResetPwdInput) (err error)
		SetStatus(ctx *gin.Context, input user_query.UserSetStatusInput) (err error)
		List(ctx *gin.Context, input user_query.UserListInput) (out user_query.UserListOut, err error)
		GetInfo(ctx *gin.Context, id int) (out user_query.UserListItem, err error)
	}
)

func RegisterNewUser(i iUser) {
	localUserService.UserService = i
}
