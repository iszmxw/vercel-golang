package mysql

import (
	"fmt"
	"goapi/pkg/config"
	MysqlLog "goapi/pkg/logger"
	"goapi/pkg/logger/zapgorm2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {
	var err error

	// 初始化 MySQL 连接信息
	var (
		host     = config.GetString("database.mysql.host")
		port     = config.GetString("database.mysql.port")
		database = config.GetString("database.mysql.database")
		username = config.GetString("database.mysql.username")
		password = config.GetString("database.mysql.password")
		prefix   = config.GetString("database.mysql.prefix")
		charset  = config.GetString("database.mysql.charset")
	)
	tls := "true"
	if len(os.Getenv("TLS")) > 0 {
		tls = os.Getenv("TLS")
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=%s&interpolateParams=true&charset=%s&parseTime=%t&loc=%s", username, password, host, port, database, tls, charset, true, "Local")
	gormConfig := mysql.New(mysql.Config{
		DSN: dns,
	})
	// 追踪mysql日志
	logger := zapgorm2.New(MysqlLog.Logger)
	logger.SetAsDefault()
	// 准备数据库连接池
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix, // 表名前缀，`User` 的表名应该是 `go_users`
			SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `go_user`
		},
		Logger: logger,
	})
	if err != nil {
		log.Println("数据库链接失败", dns, err.Error())
	}
	err = DB.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", func(d *gorm.DB) {
		// 在未找到时引发错误
		d.Statement.RaiseErrorOnNotFound = false
	})
	if err != nil {
		log.Println("数据库 disable_raise_record_not_found 失败", dns, err.Error())
	}
	return DB
}
