package bootstrap

import (
	"github.com/gin-gonic/gin"
	"goapi/templates"
	"html/template"
	"io/fs"
	"net/http"
)

func SetupTemplate(router *gin.Engine) {
	// 将静态文件夹目录绑定到 statics 路由下访问
	defaultFads, _ := fs.Sub(templates.Statics, "statics")
	router.StaticFS("/statics", http.FS(defaultFads))
	// 调用函数列出所有文件
	_template, _ := template.ParseFS(
		templates.Default,
		"default/*.html",
		"default/**/*",
	)
	router.SetHTMLTemplate(_template)
}
