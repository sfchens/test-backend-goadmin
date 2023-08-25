package user

import (
	"csf/core/query/user_query"
	"csf/core/service"
	"fmt"
)

type sUser struct{}

func init() {
	service.RegisterNewUser(NewUserService())
}
func NewUserService() *sUser {
	return &sUser{}
}

func (s *sUser) Add(input user_query.UserAdd) {
	fmt.Printf("232323")
}
