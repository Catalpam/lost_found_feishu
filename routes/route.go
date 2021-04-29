package routes

import (
	"github.com/gin-gonic/gin"
	"lost_found/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/getstudentinfo", controller.GetStudentInfo)
	return r
}

