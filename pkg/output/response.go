package output

import (
	"goapi/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Rjson 成功返回封装 参数 data interface{} 类型为可接受任意类型
func Rjson(c *gin.Context, result interface{}, msg string, code int64) {
	reqId, _ := c.Get("Tracking-Id")
	//返回数据
	response := gin.H{
		"code":    code,
		"data":    result,
		"reqId":   reqId,
		"message": msg,
	}
	logger.Info(response)
	c.JSON(200, response)
	c.Abort()
	return
}
