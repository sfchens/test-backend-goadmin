package user

import (
	"csf/core/query/user_query"
	"csf/core/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type sUser struct{}

func init() {
	service.RegisterNewUser(NewUserService())
}
func NewUserService() *sUser {
	return &sUser{}
}

func (s *sUser) Add(ctx *gin.Context, input user_query.UserAddOrEditInput) (err error) {
	fmt.Printf("232323")
	return
}
