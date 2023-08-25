package service

import (
	"csf/core/query/user_query"
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
		Add(input user_query.UserAdd)
	}
)

func RegisterNewUser(i iUser) {
	localUserService.UserService = i
}
