package echo

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/templates"
	"log"
	"net/http"
)

// Rjson 成功返回封装 参数 data interface{} 类型为可接受任意类型
func Rjson(c *gin.Context, result interface{}, msg string, code int) {
	reqId, _ := c.Get("Tracking-Id")
	rdata := cmap.New().Items()
	rdata["reqId"] = reqId
	rdata["code"] = code
	if result != nil {
		rdata["data"] = result
	}
	if len(msg) > 0 {
		rdata["msg"] = msg
	} else {
		rdata["msg"] = "success."
	}
	c.JSON(200, rdata)
	return
}

// Error  错误返回封装
func Error(c *gin.Context, code string, lastMsg string) {
	var logInfo []zapcore.Field
	if len(logger.RequestId) > 0 {
		logInfo = append(logInfo, zap.Any("RequestId", logger.RequestId))
	}
	code, firstMsg := GetCode(code, c.Request.Header.Get("Language")) // 最终code
	fullMsg := firstMsg + lastMsg                                     // 最终msg
	logInfo = append(logInfo, zap.Any(code, fullMsg))
	logger.Logger.WithOptions(zap.AddCallerSkip(1)).Info("返回错误", logInfo...)
	Rjson(c, nil, fullMsg, helpers.StringToInt(code))
}

// Success  错误返回封装
func Success(c *gin.Context, result interface{}, msg string) {
	var logInfo []zapcore.Field
	if len(logger.RequestId) > 0 {
		logInfo = append(logInfo, zap.Any("RequestId", logger.RequestId))
	}
	logInfo = append(logInfo, zap.Any("返回数据", result))
	logger.Logger.WithOptions(zap.AddCallerSkip(1)).Info("成功返回", logInfo...)
	Rjson(c, result, msg, 1)
}

// HTHL 渲染 Jet 模板
func HTHL(c *gin.Context, templatePath string, data map[string]interface{}) {
	view, err := templates.Views.GetTemplate(templatePath)
	if err != nil {
		log.Printf("[ERROR] 模板加载错误: %v", err)
		c.String(http.StatusInternalServerError, "模板错误")
		return
	}

	vars := make(jet.VarMap)
	for key, value := range data {
		vars.Set(key, value)
	}

	err = view.Execute(c.Writer, vars, nil)
	if err != nil {
		log.Printf("[ERROR] 渲染失败: %v", err)
		c.String(http.StatusInternalServerError, "渲染错误")
	}
}
