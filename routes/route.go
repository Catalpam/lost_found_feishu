package routes

import (
	"github.com/gin-gonic/gin"
	"lost_found/controller"
	"lost_found/handler"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/getstudentinfo", controller.GetStudentInfo)
	r.POST("/login", controller.SetLoginCookie)
	//r.POST("/eventinit", controller.PrintAllData)
	r.POST("/eventinit", handler.EventHandler)

	return r
}

