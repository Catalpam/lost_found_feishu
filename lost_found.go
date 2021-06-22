package main

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/controller"
	"lost_found/controller/admin"
	"lost_found/controller/general"
	"lost_found/controller/miniController"
	"lost_found/controller/webController"
	"lost_found/handler"
	"lost_found/miniMiddleWare"
	"lost_found/webMiddleWare"
)

func main() {
	//println("Debug: Send Thx")
	//cardMessage.SendCardMessage(
	//	"ou_273dbf68377bc685de3dd11c6102f879",
	//	cardMessage.FoundClaimCard(cardMessage.FoundClaim{
	//		ItemSubtype:  "found.SubType",
	//		LeaveMessage: "found.LosterComment",
	//		ImageKey:  "img_v2_db10c690-a054-4148-a3c1-bdc2a4e075cg" ,
	//	}),
	//)

	//cardMessage.SendCardMessage(
	//	"ou_273dbf68377bc685de3dd11c6102f879",
	//	cardMessage.SameNameCard(cardMessage.SameName{
	//		FoundId:    9,
	//		FoundPlace: "二教吧",
	//		ImageKey:   "img_2e320ed6-8a53-405a-b075-edec757ff25g",
	//	}),
	//)
	//cardMessage.SendCardMessage(
	//	"ou_273dbf68377bc685de3dd11c6102f879",
	//	cardMessage.LostAddedCard(cardMessage.LostAdded{
	//		LostId: 	"strconv.Itoa(int(newLost.ID))",
	//		ItemSubtype: "newLost.TypeSubName",
	//		LostDate:    "2021-05-01"+" "+ "上午",
	//		LostPlace:   "sendPlaceStr",
	//	}),
	//)


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
	////更新数据库内容
	//common.ItemTypeInitial()
	//common.UpdateStudentList()
	//common.PlaceInitial()
	common.UpdateStudentList()
	////加载根路由组
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
	miniRoutes.POST("/getplaces", general.GetPlaces)
	miniRoutes.POST("/addfound", miniController.AddFound)
	miniRoutes.POST("/getfound", miniController.GetFoundList)
	miniRoutes.POST("/addlost", miniController.AddLost)
	miniRoutes.POST("/uploadimg", controller.UploadImg)
	miniRoutes.POST("/claim", miniController.CliamFound)
	miniRoutes.POST("/thanks", miniController.ThanksMsg)
	miniRoutes.GET("/me", miniController.GetMeInfo)
	miniRoutes.POST("/me", miniController.GetMeDetail)
	miniRoutes.POST("/foundBySelf", miniController.HasFoundBySelf)
	miniRoutes.POST("/BulkAddFound", miniController.BulkAddFound)
	miniRoutes.GET("/GetPermisson", miniController.GetPermisson)
	miniRoutes.POST("/GetPermisson", miniController.GetPermisson)






	//后台管理Web登录
	r.GET("/weblogin", webController.Login302)
	r.GET("/webcode", webController.SetCookies)
	//后台管理Web路由组
	webRoutes := r.Group("/management")
	//使用Auth中间件进行认证
	webRoutes.Use(webMiddleWare.WebAuthMiddleWare())

	webRoutes.GET("/me", webController.MeInfo)
	// 招领管理
	webRoutes.GET("/found", webController.GetFounds)
	webRoutes.GET("/lost", webController.GetLosts)
	webRoutes.GET("/match", webController.GetMatches)
	webRoutes.GET("/deletefound", webController.DelFound)
	webRoutes.GET("/deletelost", webController.DelLost)

	// 地点管理
	webRoutes.GET("/place", webController.GetPlaces)
	webRoutes.GET("/bigplace", webController.GetPlacesBig)

	webRoutes.GET("/putsubplace", webController.AddSubPlace)
	webRoutes.GET("/putplace", webController.AddPlace)

	webRoutes.GET("/deletesubplace", webController.DelSubPlace)
	webRoutes.GET("/deleteplace", webController.DelPlace)

	webRoutes.GET("/editplacebig", webController.EditPlaceBig)
	webRoutes.GET("/editplacesmall", webController.EditPlaceSmall)


	// 类型管理
	webRoutes.GET("/typesmall", webController.GetTypesSmall)
	webRoutes.GET("/typebig", webController.GetTypesBig)

	webRoutes.GET("/puttypesmall", webController.AddTypeSmall)
	webRoutes.GET("/puttypebig", webController.AddTypeBig)

	webRoutes.GET("/edittypebig", webController.EditTypeBig)
	webRoutes.GET("/edittypesmall", webController.EditTypeSmall)

	webRoutes.GET("/deletetypebig", webController.DelTypeBig)
	webRoutes.GET("/deletetypesmall", webController.DelTypeSmall)

	webRoutes.GET("/typeDefault", common.ItemTypeInitial)
	webRoutes.GET("/placeDefault", common.PlaceInitial)

	// 权限管理
	webRoutes.GET("/getPlaces", admin.GetPlaces)
	webRoutes.GET("/setPlaceAdmin", admin.SetPlaceAdmin)
	webRoutes.GET("/setAdminPrivilage", admin.SetAdminPrivilage)
	webRoutes.GET("/getPrivilages", admin.GetPrivilages)
	webRoutes.GET("/searchUser", admin.SearchUser)
	webRoutes.GET("/getAdmins", admin.GetAdmins)

	return r
}
