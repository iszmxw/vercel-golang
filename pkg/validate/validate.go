package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goapi/pkg/echo"
)

// 校验参数

func ParamsError(c *gin.Context, params interface{}) bool {
	var locale = "en"
	ValidateLang := new(echo.ValidateLang)
	ValidateLang.Validate = validator.New()
	c.Request.Header.Set("Language", locale)
	vErr := ValidateLang.SetLang(locale).Validate.Struct(params)
	if vErr != nil {
		msg := ValidateLang.Translate(vErr, params)
		echo.Error(c, "ParamsError", "："+msg)
		return true
	}
	return false
}
