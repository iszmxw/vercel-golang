package client

import (
	clientV1 "goapi/app/controllers/client/v1"
	middlewaresV1 "goapi/app/middlewares/v1"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册Client路由
func RegisterClientRoutes(router *gin.RouterGroup) {
	//router.Use(middlewaresV1.AddressLimit()) // 地区限制
	router.Use(middlewaresV1.Client())
	// 路由分组 客户端 模块
	AppRoute := router.Group("/app")
	{
		{ // V1 版本
			clientV1Group := new(clientV1.Group)
			V1Route := AppRoute.Group("/v1")

			// 设置token接口
			config := V1Route.Group("/config")
			{
				// 设置token
				config.GET("/setKlineToken.json", clientV1Group.ConfigController.SetKlineToken)
				// 检查多语言是否包含中文，然后发送警告
				config.POST("/checkLanguage.json", clientV1Group.ConfigController.CheckLanguage)
			}

			// 微信接口
			wechat := V1Route.Group("/wechat")
			{
				wechat.POST("/init.json", clientV1Group.WechatController.ServeWechat)
				wechat.POST("/userSynchronize.json", clientV1Group.WechatController.Synchronize)
				// 发送消息接口
				wechat.POST("/Send.json", clientV1Group.WechatController.Send)
			}

			// SEO优化
			seo := V1Route.Group("/seo")
			{
				seo.GET("/getGoogleIndex", clientV1Group.SeoController.GetGoogleIndex)
			}
		}

	}
}
