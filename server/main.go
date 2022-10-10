package main

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
	"time"
)

func main() {
	//初始化配置
	global.VP = initialize.InitViper()
	//初始化日志
	global.LOG = initialize.InitZap()
	global.DB = initialize.InitDB()
	if global.DB != nil {
		db, _ := global.DB.DB()
		// 程序结束前关闭数据库连接
		defer db.Close()
	}
	global.REDIS = initialize.InitRedis()
	Router := initialize.Routers()
	port := fmt.Sprintf(":%d", global.CONFIG.Server.Port)
	s := initialize.InitServer(port, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", port))
	global.LOG.Error(s.ListenAndServe().Error())
}
