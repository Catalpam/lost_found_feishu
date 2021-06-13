package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func MeInfo(ctx *gin.Context)  {
	OpenId := ctx.MustGet("open_id").(string)
	db := common.GetDB()
	var user dbModel.User
	var privilage dbModel.Privilege
	db.Where("open_id=?",OpenId).First(&user)
	db.Where("open_id=?",OpenId).First(&privilage)
	if !(privilage.Permission == "2" || privilage.Permission == "3") {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 403,
			"avatar":user.Avatar,
			"name":user.Name,
			"privilageCode": privilage.Permission,
			"privilageName": permissionCode2String(privilage.Permission),
			"msg":  "你的权限不足，请向相关管理人员申请使用权限！",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"avatar":user.Avatar,
		"name":user.Name,
		"privilageCode": privilage.Permission,
		"privilageName": permissionCode2String(privilage.Permission),
		"msg":  "个人信息获取成功",
	})
}

func permissionCode2String(code string)string {
	if code == "0" {return "普通用户"}
	if code == "1" {return "地点管理员"}
	if code == "2" {return "平台审核管理员"}
	if code == "3" {return "超级管理员"}
	return ""
}

