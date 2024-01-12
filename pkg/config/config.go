package config

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"os"
)

// Viper Viper 库实例
var Viper *viper.Viper

// StrMap 简写 —— map[string]interface{}
type StrMap map[string]interface{}

// init() 函数在 import 的时候立刻被加载
func init() {
	// 加载环境变量配置文件
	_ = godotenv.Load()
	// 1、从环境变量读取配置文件
	jsonConfig := os.Getenv("CONFIG")
	fmt.Println("jsonConfig", jsonConfig)
	// 2、 初始化 Viper 库
	Viper = viper.New()
	Viper.SetConfigType("json") // 设置配置文件的类型
	// 3、创建io.Reader
	err := Viper.ReadConfig(bytes.NewBuffer([]byte(jsonConfig)))
	if err != nil {
		panic(err)
	}
	// 4. 设置环境变量前缀，用以区分 Go 的系统环境变量
	Viper.SetEnvPrefix("appenv")
	// 5. Viper.Get() 时，优先读取环境变量
	Viper.AutomaticEnv()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

// Add 新增配置项
func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

// Get 获取配置项，允许使用点式获取，如：app.name
func Get(path string, defaultValue ...interface{}) interface{} {
	// 不存在的情况
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

// GetStringSlice 获取 String 类型的配置信息
func GetStringSlice(path string, defaultValue ...interface{}) []string {
	return cast.ToStringSlice(Get(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}
