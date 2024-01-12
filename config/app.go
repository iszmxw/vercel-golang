package config

import (
	"goapi/pkg/config"
)

// app 参数
func init() {
	config.Add("app", config.StrMap{

		// 应用名称，暂时没有使用到
		"name": config.Env("APP_NAME", "goapi"),

		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "production"),

		// 是否进入调试模式
		"debug": config.Env("APP_DEBUG", false),
		// 是否进入调试模式
		"pprof": config.Env("APP_PPROF", false),

		// 应用服务端口
		"port": config.Env("APP_PORT", "8888"),

		// 系统加密中加密数据时使用
		"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

		// 用以生成链接
		"url": config.Env("APP_URL", "http://127.0.0.1:8888"),

		// php服务的地址
		"php_url": config.Env("PHP_URL", "http://127.0.0.1:10087"),

		// IMG URL用于存放图片的路径
		"img_url": config.Env("IMG_URL", "http://localhost:80"),

		"https": config.Env("HTTPS", "0"),

		// IP地区限制
		"address_limit": config.Env("ADDRESS_LIMIT", false),
	})
}
