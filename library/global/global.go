package global

import "github.com/gin-gonic/gin"

var (
	GinEngine  *gin.Engine
	RouterList = make(map[string][]func(v1 *gin.RouterGroup))
)
