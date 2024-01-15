package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/middlewares/common"
	conf "goapi/pkg/config"
	"goapi/routes/client"
	"goapi/routes/web"
	"net/http"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) *gin.Engine {
	router.Use(common.TraceLogger()) // 日志追踪
	router.Use(common.Cors())        // 跨域
	router.NoRoute(NoResponse)
	router.GET("/route", func(context *gin.Context) {
		requestId, _ := context.Get("Tracking-Id")
		context.String(200, fmt.Sprintf("Hello World!：%v\n\n", requestId))
		context.String(200, fmt.Sprintf("当前的redis地址是：%s\n\n\n\t", conf.GetString("redis.host")))
		context.String(200, "下面是所有接口服务：\n\n\n\t")
		routers := router.Routes()
		for _, v := range routers {
			context.String(200, fmt.Sprintf("Method：%v\tURL：%v  \n\tHandler: %v\n\t=====================\n\t\n\t", v.Method, v.Path, v.Handler))
		}
	})
	// client 接口
	client.RegisterClientRoutes(router.Group("/"))
	// web 相关页面
	web.RegisterWebRoutes(router.Group("/"))
	return router
}

func NoResponse(c *gin.Context) {
	//返回 404 状态码
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404, page not exists!",
	})
}
