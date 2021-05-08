package main

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/controller"
	"lost_found/controller/webController"
	"lost_found/handler"
)

func main() {
	//数据库加载初始化
	db := common.InitDB()
	defer db.Close()

	//更新数据库内容
	common.ThingDbDefaultInit()
	common.UpdateStudentList()

	//加载根路由组
	r := gin.Default()
	r = CollectRoute(r)
	r.Run(":1111")
}

func CollectRoute(r *gin.Engine) *gin.Engine {
	//️处理飞书服务器传来的事件
	r.POST("/webhook", handler.EventHandler)

	//小程序路由组
	//r.POST("/miniLogin", webController.SetLoginStatus)
	miniRoutes := r.Group("/miniapp")
	miniRoutes.POST("/userinfo", controller.GetUserInfo)
	miniRoutes.POST("/gettypes", controller.GetTypes)
	miniRoutes.POST("/addfound", controller.AddFound)



	//后台管理Web登录
	r.GET("/weblogin", webController.Login302)
	r.GET("/webcode", webController.SetCookies)
	//后台管理Web路由组
	webRoutes := r.Group("/web")
	//webRoutes.Use(middleware.AuthMiddleware())
	webRoutes.POST("/userinfo", controller.GetUserInfo)
	webRoutes.POST("/addfound", controller.AddFound)

	return r
}
