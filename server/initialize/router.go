package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"server/global"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	//systemRouter := router.RouterGroupApp.System
	//exampleRouter := router.RouterGroupApp.Example

	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		//systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	//PrivateGroup := Router.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//{
	//	systemRouter.InitApiRouter(PrivateGroup) // 注册功能api路由
	//
	//}

	//InstallPlugin(Router) // 安装插件

	global.LOG.Info("router register success")
	return Router
}
