package utils

import (
	"csf/library/easy_config"
	"csf/library/easy_session"
	"csf/library/global"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserInfo struct {
	Id        uint   `json:"id" description:"ID"`
	Username  string `json:"username" description:"账号"`
	Realname  string `json:"realname" description:"真名"`
	Email     string `json:"email" description:"邮箱"`
	Phone     string `json:"phone" description:"号码"`
	LoginType string `json:"login_type" description:"登录类型"`
}

// GetUserName 获取用户账号
func GetUserName(ctx *gin.Context) (username string) {
	userInfo, err := GetUserInfo(ctx)
	if err != nil {
		return
	}
	if userInfo.Id > 0 {
		username = userInfo.Username
		return
	}
	return
}

// GetUserInfo 获取用户信息
func GetUserInfo(ctx *gin.Context) (userinfo UserInfo, err error) {
	userInfoJson := easy_session.NewCustomSession(ctx).Get(global.UserInfoKey)
	val, ok := userInfoJson.(string)
	if !ok {
		err = errors.New("获取session用户信息失败")
		return
	}
	err = json.Unmarshal([]byte(val), &userinfo)
	if err != nil {
		return
	}
	return
}

// GetToken 获取token
func GetToken(ctx *gin.Context) (token string) {
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

func GetModulesName(ctx *gin.Context) (module string) {

	path := ctx.Request.URL.Path
	strArr := strings.Split(path, "/")

	if len(strArr) > 0 {
		module = strArr[1]
	} else {
		module = easy_config.Config.App.Name
	}
	return
}
