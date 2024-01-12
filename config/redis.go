package config

import (
	"fmt"
	"goapi/pkg/config"
)

// redis 参数
func init() {
	fmt.Println("初始化redis配置")
	config.Add("redis", config.StrMap{
		"host":     config.Env("REDIS_HOST", "127.0.0.1"),
		"port":     config.Env("REDIS_PORT", 6379),
		"password": config.Env("REDIS_PASSWORD", ""),
		"db":       config.Env("REDIS_DB", 0),
	})
}
