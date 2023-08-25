package user

import (
	"csf/app/admin/request/user_req"
	"csf/app/admin/service/user_service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cUserApi struct {
}

func NewUserApi() *cUserApi {
	return &cUserApi{}
}

func (c *cUserApi) Add(ctx *gin.Context) {
	var (
		err error

		req user_req.UserAddOrEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = user_service.NewUserService(ctx).Add(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

func (c *cUserApi) List(ctx *gin.Context) {

}
