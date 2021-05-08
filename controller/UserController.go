package controller

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/dbModel"
	contact "lost_found/service/contact/v3"
	"net/http"
)

var contactService = contact.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func GetUserInfo(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	OpenId, ok_1 := ctx.GetPostForm("open_id")
	Name, ok_2 := ctx.GetPostForm("name")

	//检测参数是否合法
	if ok_1 != true && ok_2 != true {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": nil,
			"msg": "缺少参数, 需要open_id或者name",
		})
		return
	}

	//定义数据库映射模型 User
	var student dbModel.User

	//当参数为open_id时的响应 Student初始化
	if ok_1 == true {
		DB.Where("open_id = ?", OpenId).First(&student)
		if student.ID == 0{
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"data": nil,
				"msg": "该open_id所请求的用户不存在",
			})
			return
		}
	} else {
		//当参数为name时 Student初始化
		if ok_2 == true  {
			DB.Where("name = ?", Name).First(&student)
			if student.ID == 0{
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"code": 422,
					"data": nil,
					"msg": "该name所请求的用户不存在",
				})
				return
			}
		}
	}

	//查询成功后返回的请求
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"open_id": student.OpenId,
			"name": student.Name,
			"department_id": student.DepartmentId,
			"avatar": student.Avatar,
			//"mobile": student.Mobile,
		},
		"msg": "请求成功!",
	})
}

