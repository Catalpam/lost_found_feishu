package main

import (
	"github.com/gin-gonic/gin"
	"lost_found/cardMessage"
	"lost_found/common"
	"lost_found/controller"
	"lost_found/controller/general"
	"lost_found/controller/miniController"
	"lost_found/controller/webController"
	"lost_found/handler"
	"lost_found/miniMiddleWare"
)

func main() {

	cardMessage.SendCardMessage(
		"ou_273dbf68377bc685de3dd11c6102f879",
		cardMessage.ReSurveyCard(cardMessage.ReSurvey{
			LostId:      "0",
			ItemSubtype: "一个月后拿到的MBP",
			FoundDate:   "2222-10-1",
			ImageKey:    "img_2e320ed6-8a53-405a-b075-edec757ff25g",
		}),
	)


	//cardMessage.SendCardMessage(
	//	"ou_273dbf68377bc685de3dd11c6102f879",
	//	cardMessage.FoundClaimCard(cardMessage.FoundClaim{
	//		ItemSubtype:  "一个月后拿到的MBP",
	//		LeaveMessage: "来自失主的感谢",
	//		ImageKey:     "img_2e320ed6-8a53-405a-b075-edec757ff25g",
	//	}),
	//)
	//
	//cardMessage.SendCardMessage(
	//	"ou_273dbf68377bc685de3dd11c6102f879",
	//	cardMessage.ThanksHasSendCard(cardMessage.ThanksHasSend{
	//		ItemSubtype: "一个月后拿到的MBP",
	//		FoundDate:   "2021-5-11",
	//		ImageKey:    "img_2e320ed6-8a53-405a-b075-edec757ff25g",
	//	}),
	//)
	//
	//cardMessage.SendCardMessage(
	//	"ou_273dbf68377bc685de3dd11c6102f879",
	//	cardMessage.SendUser2FounderCard(cardMessage.SendUser2Founder{
	//		ItemSubtype: "一个月后拿到的MBP",
	//		FoundDate:   "2021-5-11",
	//		ImageKey:    "img_2e320ed6-8a53-405a-b075-edec757ff25g",
	//	}),
	//)
	//
	//cardMessage.SendCardMessage(
	//	"ou_273dbf68377bc685de3dd11c6102f879",
	//	cardMessage.SendUser2LosterCard(cardMessage.SendUser2Loster{
	//		FounderName: "鬼才冯梓耕",
	//		ItemSubtype: "一个月后拿到的MBP",
	//		FoundDate:   "2021-5-11",
	//		ImageKey:    "img_2e320ed6-8a53-405a-b075-edec757ff25g",
	//	}),
	//)

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
	panic(r.Run(":1111"))
}

func CollectRoute(r *gin.Engine) *gin.Engine {
	//通用
	r.GET("/image", controller.GetImage)

	//️处理飞书服务器传来的事件
	r.POST("/webhook", handler.EventHandler)
	r.POST("/card", handler.CardEvent())

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
	miniRoutes.POST("/getfound", miniController.GetFoundList)
	miniRoutes.POST("/addlost", miniController.AddLost)
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
