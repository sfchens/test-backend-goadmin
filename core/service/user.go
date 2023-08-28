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
		Add(ctx *gin.Context, input user_query.UserAddOrEditInput) (err error)
	}
)

func RegisterNewUser(i iUser) {
	localUserService.UserService = i
}
