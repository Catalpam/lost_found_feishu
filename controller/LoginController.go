package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetLoginCookie(ctx *gin.Context) {

	ctx.SetCookie("site_cookie", "cookievalue", 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "",
		"msg": "请求成功!",
	})

}
