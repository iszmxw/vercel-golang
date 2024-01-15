package web

import (
	"github.com/gin-gonic/gin"
	v1 "goapi/app/controllers/web/v1"
)

// RegisterWebRoutes 注册 Web 路由
func RegisterWebRoutes(router *gin.RouterGroup) {
	GroupV1 := new(v1.GroupV1)
	webV1 := router.Group("/")
	{
		webV1.GET("/", GroupV1.IndexController.Index)
		webV1.GET("/home", GroupV1.IndexController.Home)
		webV1.GET("/view", GroupV1.IndexController.View)
	}
}
