package echo

import (
	"fmt"
	"github.com/go-playground/locales"
	locEn "github.com/go-playground/locales/en"
	locZh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/pkg/echo/code/en"
	"goapi/pkg/echo/code/insa"
	"goapi/pkg/echo/code/zh"
	"goapi/pkg/logger"
	"reflect"
	"strings"
)

type ValidateLang struct {
	Lang       string
	uni        *ut.UniversalTranslator
	importLang locales.Translator
	trans      ut.Translator
	Validate   *validator.Validate
}

func GetCode(code, Language string) (string, string) {
	var t reflect.Type
	switch strings.ToLower(Language) {
	case "zh":
		t = reflect.TypeOf(&zh.Code{}).Elem()
		break
	case "en":
		t = reflect.TypeOf(&en.Code{}).Elem()
		break
	case "insa":
		t = reflect.TypeOf(&insa.Code{}).Elem()
		break
	default: // 默认中文
		t = reflect.TypeOf(&zh.Code{}).Elem()
		break
	}
	field, ok := t.FieldByName(code)
	if !ok {
		return "0", "unknown mistake"
	}
	return field.Tag.Get("code"), field.Tag.Get("msg")
}

func (V *ValidateLang) SetLang(locale string) *ValidateLang {
	logger.Info(locale)
	var err error
	V.Lang = locale
	switch V.Lang {
	case "zh":
		V.importLang = locZh.New()
		V.uni = ut.New(V.importLang, V.importLang)
		V.trans, _ = V.uni.GetTranslator(V.Lang)
		err = zhTranslations.RegisterDefaultTranslations(V.Validate, V.trans)
		fmt.Println(err)
		break
	case "en":
		V.importLang = locEn.New()
		V.uni = ut.New(V.importLang, V.importLang)
		V.trans, _ = V.uni.GetTranslator(V.Lang)
		err = enTranslations.RegisterDefaultTranslations(V.Validate, V.trans)
		break
	default: // 默认中文
		V.importLang = locZh.New()
		V.uni = ut.New(V.importLang, V.importLang)
		V.trans, _ = V.uni.GetTranslator(V.Lang)
		err = zhTranslations.RegisterDefaultTranslations(V.Validate, V.trans)
		break
	}
	if err != nil {
		fmt.Println("validator 翻译出错")
	}
	return V
}

// Translate 翻译工具
func (V *ValidateLang) Translate(err error, s interface{}) string {
	result := cmap.New()
	t := reflect.TypeOf(s)
	if reflect.TypeOf(err).String() == "*validator.InvalidValidationError" {
		return err.(*validator.InvalidValidationError).Error()
	}
	for _, errs := range err.(validator.ValidationErrors) {
		// 使用反射方法获取struct种的json标签作为key --重点2
		var k string
		if field, ok := t.FieldByName(errs.StructField()); ok {
			k = field.Tag.Get("json")
		}
		if k == "" {
			k = errs.StructField()
		}
		result.Set(k, errs.Translate(V.trans))
	}
	return getFirstMessage(result.Items())
}

func getFirstMessage(messages map[string]interface{}) string {
	for _, val := range messages {
		return val.(string)
	}
	return ""
}
