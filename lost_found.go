package main

import (
	"lost_found/common"
	"lost_found/controller"
	"lost_found/controller/general"
	"lost_found/controller/miniController"
	"lost_found/controller/webController"
	"lost_found/handler"
	"lost_found/miniMiddleWare"

	"github.com/gin-gonic/gin"
)

func main() {
	//Redis缓存加载初始化
	common.RedisInit()
	//数据库加载初始化
	db := common.InitDB()
	defer db.Close()

	//更新数据库内容
	common.ItemTypeInitial()
	common.UpdateStudentList()
	common.PlaceInitial()

	//加载根路由组
	r := gin.Default()
	r = CollectRoute(r)
	r.Run(":1111")
}

func CollectRoute(r *gin.Engine) *gin.Engine {
	//通用
	r.GET("/image", controller.GetImage)

	//️处理飞书服务器传来的事件
	r.POST("/webhook", handler.EventHandler)

	//小程序登录
	r.POST("/minilogin", miniController.SetCookies)
	//小程序路由组初始化
	miniRoutes := r.Group("/miniapp")
	//使用Auth中间件进行认证
	miniRoutes.Use(miniMiddleWare.MiniAuthMiddleWare())
	//小程序路由组
	miniRoutes.POST("/userinfo", controller.GetUserInfo)
	miniRoutes.POST("/gettypes", general.GetTypes)
	miniRoutes.POST("/getcampus", general.GetCampus)
	miniRoutes.POST("/getplaces", general.GetPlaces)
	miniRoutes.POST("/addfound", miniController.AddFound)
	miniRoutes.POST("/getfound", miniController.GetFound)
	miniRoutes.POST("/uploadimg", controller.UploadImg)

	//后台管理Web登录
	r.GET("/weblogin", webController.Login302)
	r.GET("/webcode", webController.SetCookies)
	//后台管理Web路由组
	webRoutes := r.Group("/web")
	//webRoutes.Use(middleware.AuthMiddleware())
	webRoutes.POST("/userinfo", controller.GetUserInfo)
	webRoutes.POST("/addfound", miniController.AddFound)

	return r
}