package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", *AppPort),
		Handler: router,
	}

	// 在 goroutine 中启动服务器
	go func() {
		// 启动路由
		logger.Info("启动路由")
		logger.Info(fmt.Sprintf("当前环境:%v", *AppPort))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("服务器启动失败:", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("正在关闭服务器...")

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务器关闭失败:", err)
	}

	logger.Info("服务器已退出")
}
