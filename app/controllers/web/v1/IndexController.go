package v1

import (
	"github.com/gin-gonic/gin"
	"goapi/pkg/echo"
	"goapi/pkg/portscanner"
	"time"
)

type IndexController struct {
	BaseController
}

func (h *IndexController) Test(c *gin.Context) {
	ip := c.Query("ip")
	PortRange := c.Query("PortRange")
	if len(ip) == 0 {
		echo.Error(c, "ParamsLost", "ip 参数不能为空")
		return
	}
	scanner := portscanner.Scanner{
		//PortRange:   "12111-12120", // 端口范围
		IP:          ip, // 目标IP
		Timeout:     300 * time.Millisecond,
		Concurrency: 100, // 并发数
	}
	if len(PortRange) > 0 {
		scanner.PortRange = PortRange
	}
	results := scanner.Run()
	echo.Success(c, results, "端口扫描成功")
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
