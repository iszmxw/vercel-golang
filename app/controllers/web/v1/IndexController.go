package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	BaseController
}

func (h *IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "World!",
	})
}
