package utils

import (
	"csf/common/mysql/model"
	"csf/library/custom_session"
	"csf/library/global"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserInfo struct {
	Id       uint   `json:"id" description:"ID"`
	Username string `json:"username" description:"账号"`
	Realname string `json:"realname" description:"真名"`
	Email    string `json:"email" description:"邮箱"`
	Phone    string `json:"phone" description:"号码"`
}

// GetUserName 获取用户账号
func GetUserName(ctx *gin.Context) (username string) {
	userInfo := GetUserInfo(ctx)
	if userInfo.Id > 0 {
		username = userInfo.Username
		return
	}
	token := GetAuthorization(ctx)
	if token == "" {
		return ""
	}
	//mc, err := my_jwt.NewJWT().ParseToken(token)
	//if err != nil {
	//	return ""
	//}
	//username = mc.BaseClaims.Username
	return
}
func GetUserId(ctx *gin.Context) (id uint) {
	//token := GetAuthorization(ctx)
	//if token == "" {
	//	return 0
	//}
	//mc, err := my_jwt.NewJWT().ParseToken(token)
	//if err != nil {
	//	return 0
	//}
	//return uint(mc.BaseClaims.Id)
	return
}

// GetAdminInfo 获取管理员登录信息
func GetAdminInfo(ctx *gin.Context) (data model.SysAdmin) {
	loginKey := custom_session.NewCustomSession(ctx).Get(global.LoginTypeKey)
	key, ok := loginKey.(string)
	if !ok {
		return
	}
	if key != global.LoginTypeAdmin {
		return
	}

	adminInfoTmp := custom_session.NewCustomSession(ctx).Get(global.UserInfoKey)
	val, ok := adminInfoTmp.(string)
	if !ok {
		return
	}
	_ = json.Unmarshal([]byte(val), &data)
	return
}

// GetFrontInfo 前端用户登录信息
func GetFrontInfo(ctx *gin.Context) (data model.SysUser) {
	loginKey := custom_session.NewCustomSession(ctx).Get(global.LoginTypeKey)
	key, ok := loginKey.(string)
	if !ok {
		return
	}

	if key != global.LoginTypeFront {
		return
	}

	adminInfoTmp := custom_session.NewCustomSession(ctx).Get(global.UserInfoKey)
	val, ok := adminInfoTmp.(model.SysUser)
	if !ok {
		return
	}
	data = val
	return
}

// GetUserInfo 前后端信息合并
func GetUserInfo(ctx *gin.Context) (userInfo UserInfo) {
	loginKeyTmp := custom_session.NewCustomSession(ctx).Get(global.LoginTypeKey)
	loginKey, ok := loginKeyTmp.(string)
	if !ok {
		return
	}
	if loginKey == global.LoginTypeAdmin {
		userInfoTmp := GetAdminInfo(ctx)
		userInfo = UserInfo{
			Id:       userInfoTmp.ID,
			Username: userInfoTmp.Username,
			Realname: userInfoTmp.Realname,
			Email:    userInfoTmp.Email,
			Phone:    userInfoTmp.Phone,
		}
	} else if loginKey == global.LoginTypeFront {
		userInfoTmp := GetFrontInfo(ctx)
		userInfo = UserInfo{
			Id:       userInfoTmp.ID,
			Username: userInfoTmp.Username,
			Realname: userInfoTmp.Realname,
			Email:    userInfoTmp.Email,
			Phone:    userInfoTmp.Phone,
		}
	}
	return
}

func GetAuthorization(ctx *gin.Context) (token string) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return
	}
	token = parts[1]
	return
}
