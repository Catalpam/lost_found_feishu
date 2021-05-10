package main

import (
	"lost_found/common"
	"lost_found/controller"
	"lost_found/controller/general"
	"lost_found/controller/miniController"
	"lost_found/controller/webController"
	"lost_found/handler"

	"github.com/gin-gonic/gin"
)

func main() {
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

	//小程序路由组
	r.POST("/minilogin", miniController.SetCookies)
	miniRoutes := r.Group("/miniapp")
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

//type Locations struct {
//	ShaCampus struct {
//		Building  []string `yaml:"教学楼,flow"`
//		Sushe  []string `yaml:"宿舍区,flow"`
//		ShiTang   []string `yaml:"食堂,flow"`
//		Sports []string `yaml:"运动场所,flow"`
//		Other   []string `yaml:"其他,flow"`
//	} `yaml:"沙河校区"`
//
//	QinCampus struct {
//		Building    []string `yaml:"教学楼,flow"`
//		Liu	 []string `yaml:"留学生宿舍区,flow"`
//		Shuo    []string `yaml:"硕丰苑,flow"`
//		Xue    []string `yaml:"学知苑,flow"`
//		Boshi    []string `yaml:"博翰苑,flow"`
//		ShiTang     []string `yaml:"餐厅,flow"`
//		Other     []string `yaml:"其他,flow"`
//	} `yaml:"清水河校区"`
//}

//data, err := ioutil.ReadFile("./location_config.yaml")
//if err != nil {
//println(err)
//println("读取文件失败")
//}
//
//var bs =  Locations{}
//if err := yaml.Unmarshal(data, &bs); err != nil {
//println(err)
//println("反序列化失败")
//}
//println(bs.QinCampus.Building)
