package config

import (
	"goapi/pkg/config"
)

// redis 参数
func init() {
	config.Add("wechat", config.StrMap{
		"botUrl":   config.Env("wechat.botUrl", ""),
		"guid":     config.Env("wechat.guid", ""),
		"BelongWx": config.Env("wechat.BelongWx", ""),
	})
}
