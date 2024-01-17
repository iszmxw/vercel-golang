package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init()
}

func main() {
	AppPort := flag.Int64("APP_PORT", conf.GetInt64("app.port"), "服务端口")
	flag.Parse()
	// 初始化 SQL
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	// 初始化路由绑定
	logger.Info("加载 client 路由")
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	bootstrap.SetupTemplate(app)
	router := bootstrap.SetupRoute(app)
	pprof.Register(router) // 开启 pprof
	// 启动路由
	logger.Info("启动路由")
	logger.Info(fmt.Sprintf("当前环境:%v", *AppPort))
	_ = router.Run(fmt.Sprintf(":%v", *AppPort))
}
