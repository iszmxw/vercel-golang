package handler

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"net/http"
	"strings"
	"time"
)

var (
	app *gin.Engine
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init()
	// 初始化 SQL
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	bootstrap.SetupRedis(0)
	app = gin.New()
	bootstrap.SetupTemplate(app)
	// Handling routing errors
	app.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range app.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})
	// 初始化路由绑定
	logger.Info("加载 client 路由")
	router := bootstrap.SetupRoute(app)
	pprof.Register(router) // 开启 pprof
}

/**

vercel 入口

所有的路由都从里进入，相当于平时的 main 函数
*/

func Handler(w http.ResponseWriter, r *http.Request) {
	// 启动路由
	logger.Info("启动路由")
	app.ServeHTTP(w, r)
}
