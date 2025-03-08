package v1

import (
	"github.com/gin-gonic/gin"
	"goapi/pkg/echo"
)

type IndexController struct {
	BaseController
}

func (h *IndexController) Index(c *gin.Context) {
	data := map[string]interface{}{
		"title":   "首页",
		"message": "Hello, vercel-golang，Lest go!\n框架已经准备好，开始创作吧\n欢迎来到 Gin + Jet 的主页",
		"data": map[string]interface{}{
			"keywords": "关键词",
			"desc":     "描述",
		},
	}
	data["static_url"] = "/statics/css/common.css"
	echo.HTHL(c, "pages/home.html", data)
}

func (h *IndexController) About(c *gin.Context) {
	data := map[string]interface{}{
		"title":   "关于我们",
		"message": "欢迎来到 Gin + Jet 的关于我们",
		"data": map[string]interface{}{
			"keywords": "关于我们-关键词",
			"desc":     "关于我们-描述",
		},
	}
	data["static_url"] = "/statics/css/common.css"
	echo.HTHL(c, "pages/about.html", data)
}

func (h *IndexController) View(c *gin.Context) {
	data := map[string]interface{}{
		"title":   "view 页面",
		"message": "欢迎来到 Gin + Jet 的 view 页面",
		"data": map[string]interface{}{
			"keywords": "view 页面-关键词",
			"desc":     "view 页面-描述",
		},
	}
	data["static_url"] = "/statics/css/common.css"
	echo.HTHL(c, "v/view.html", data)
}
