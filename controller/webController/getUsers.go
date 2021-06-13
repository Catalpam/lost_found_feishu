package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetUserList(ctx *gin.Context) {
	db := common.GetDB()

	//定义数据库映射模型 User
	var user dbModel.User

	//获取参数
	Name := ctx.Query("name")
	id := ctx.Query("id")

	//检测参数
	if id != "" {
		db.Where("student_id = ?", id).Find(&user)
	} else if Name != "" {
		db.Where("name = ?", Name).Find(&user)
	} else {
		db.Where("name = ?", Name).Find(&user)
	}

	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": nil,
			"msg":  "该name所请求的用户不存在",
		})
		return
	}

	//查询成功后返回的请求
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"open_id":       user.OpenId,
			"name":          user.Name,
			"department_id": user.DepartmentId,
			"avatar":        user.Avatar,
			//"mobile": user.Mobile,
		},
		"msg": "请求成功!",
	})
}

