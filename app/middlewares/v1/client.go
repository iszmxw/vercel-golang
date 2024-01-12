package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/response"
	"goapi/pkg/address"
	"goapi/pkg/config"
	"goapi/pkg/echo"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"goapi/pkg/request"
	"time"
)

// AddressLimit 地区限制
func AddressLimit() gin.HandlerFunc {
	/** 您所在的地区限制使用 */
	return func(c *gin.Context) {
		logger.Info(c.ClientIP())
		switch c.FullPath() {
		case "/app/v1/wechat/userSynchronize.json":
		case "/app/v1/wechat/Send.json":
		case "/app/v1/wechat/init.json":
		case "/app/v1/seo/getGoogleIndex":
			break
		default:
			// 正式环境检测
			if config.GetBool("app.address_limit") && address.LimitAddress(c.ClientIP()) {
				echo.Error(c, "CountryLimitUse", "")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// Client 定义客户端中间件
func Client() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.IsWebsocket() {
			// 心跳分分钟过期
			HeartBeatingKey := fmt.Sprintf("ws_index:%v:heart_beating_key", c.GetHeader("Sec-WebSocket-Key"))
			_, _ = redis.Client.Add(HeartBeatingKey, time.Now().Format("2006-01-02 15:04:05"), 300) // 5分钟过期
		}
		// 设置语言，默认英文
		SetLang(c, "en")
		path := c.FullPath()
		switch path {
		case "/app/v1/wechat/userSynchronize.json":
		case "/app/v1/config/setKlineToken.json":
		case "/app/v1/wechat/Send.json":
		case "/app/v1/wechat/init.json":
		case "/app/v1/seo/getGoogleIndex":
			// 继续往下面执行
			c.Next()
			break
		default:
			CheckLogin(c)
			c.Next()
			break
		}
	}
}

// CheckLogin 检测登录
func CheckLogin(c *gin.Context) {
	// 获取 "token"
	tokenString := request.GetParam(c, "token")
	if len(tokenString) <= 0 || tokenString == "<nil>" {
		if !c.IsWebsocket() {
			logger.Error(errors.New("未检测到token"))
			echo.Error(c, "LoginInvalid", "")
			c.Abort()
			return
		}
	}
	strInfo, _ := redis.Client.Get(fmt.Sprintf("Client:Token:%s", tokenString))
	if len(strInfo) > 0 {
		// json字符串数组,转换成切片
		var user response.ClientUserBeans
		multiErr := json.Unmarshal([]byte(strInfo), &user)
		if multiErr != nil {
			logger.Error(errors.New("转换出错"))
			logger.Error(multiErr)
			return
		}
		logger.Info(user)
		// 保存用户到 上下文
		c.Set("user", user)
		// 是否登录
		c.Set("NoLogin", false)
	} else {
		if c.IsWebsocket() {
			c.Set("NoLogin", true)
			// 继续往下面执行
			c.Next()
		} else {
			echo.Error(c, "LoginInvalid", "")
			c.Abort()
			return
		}
	}
	// 再次确认设置语言
	SetLang(c, "en")
	// 继续往下面执行
	c.Next()
}

// SetLang 设置语言
func SetLang(c *gin.Context, lang string) {
	language := c.Request.Header.Get("Language")
	if len(language) <= 0 {
		// 默认中文
		c.Request.Header.Set("Language", lang)
	}
	l, _ := c.Get("language")
	switch l {
	case 1:
		// 中文
		c.Request.Header.Set("Language", "zh")
		break
	case 2:
		// 英文
		c.Request.Header.Set("Language", "en")
		break
	case 3:
		// 日语
		c.Request.Header.Set("Language", "jp")
		break
	}
}
