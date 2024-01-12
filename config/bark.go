package config

import (
	"goapi/pkg/config"
)

// redis 参数
func init() {
	config.Add("bark", config.StrMap{
		"url":        config.Env("bark.url", ""),
		"key.mac":    config.Env("bark.key.mac", ""),
		"key.iphone": config.Env("bark.key.iphone", ""),
		"logo":       config.Env("bark.logo", ""),
	})
}
