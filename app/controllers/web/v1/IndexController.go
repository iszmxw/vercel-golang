package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	BaseController
}

func (h *IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "vercel-golang，Lest go",
	})
}

func (h *IndexController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "我是 Home 页面，准备开始，Lest go!",
	})
}

func (h *IndexController) View(c *gin.Context) {
	c.HTML(http.StatusOK, "v/view.html", gin.H{
		"title": "View",
	})
}
