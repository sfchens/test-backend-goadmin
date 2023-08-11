package global

import "github.com/gin-gonic/gin"

var (
	GinEngine  *gin.Engine
	RouterList = make([]func(v1 *gin.RouterGroup), 0)
)
