package global

import "github.com/gin-gonic/gin"

var (
	RouterNoCheckRole = make([]func(*gin.RouterGroup), 0)
	RouterCheckRole   = make([]func(v1 *gin.RouterGroup), 0)
	GinEngine         *gin.Engine
	RouterList        = make([]func(v1 *gin.RouterGroup), 0)
)
