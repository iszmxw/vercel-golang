package config

import (
	"goapi/pkg/abnormal"
)

// Initialize 配置信息初始化
func Initialize() {
	// 触发加载本目录下其他文件中的 init 方法
	defer abnormal.Stack("项目启动", nil)
}
