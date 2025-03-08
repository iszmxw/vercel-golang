package bootstrap

import (
	"github.com/gin-gonic/gin"
	"goapi/templates"
	"io/fs"
	"log"
	"net/http"
)

func SetupTemplate(router *gin.Engine) {
	// 将静态文件夹目录绑定到 statics 路由下访问
	defaultFads, err := fs.Sub(templates.Statics, "statics")
	if err != nil {
		log.Fatalf("静态文件加载失败: %v", err)
	}
	router.StaticFS("/statics", http.FS(defaultFads))
	err = templates.LoadTemplates("views")
	if err != nil {
		log.Fatalf("模板加载失败: %v", err)
	}
}
