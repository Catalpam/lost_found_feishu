package controller

import (
	"context"
	"fmt"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	contact "lost_found/service/contact/v3"
	"lost_found/common"
	"lost_found/dbModel"
	"lost_found/core/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

var contactService = contact.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func UpdateStudentList() {
	DB := common.GetDB()
	coreCtx := core.WrapContext(context.Background())
	reqCall := contactService.Users.List(coreCtx)
	result, err := reqCall.Do()
	if err != nil {
		e := err.(*response.Error)
		fmt.Printf(tools.Prettify(e))
		return
	}

	for _, value := range result.Items {
		var dBstudent dbModel.Student
		open_id := value.OpenId
		DB.Where("open_id = ?", open_id).First(&dBstudent)
		if dBstudent.ID != 0{
			if dBstudent.Name != value.Name {
				dBstudent.Name = value.Name
			}
			if dBstudent.StudentId != value.CustomAttrs[0].Value.Text {
				dBstudent.StudentId = value.CustomAttrs[0].Value.Text
			}
			if dBstudent.Mobile != value.Mobile {
				dBstudent.Mobile = value.Mobile
			}
			if dBstudent.DepartmentId != value.DepartmentIds[0] {
				dBstudent.DepartmentId = value.DepartmentIds[0]
			}
			if dBstudent.Avatar != value.Avatar.AvatarOrigin {
				dBstudent.Avatar = value.Avatar.AvatarOrigin
			}
			continue
		}
		newStudent := dbModel.Student{
			Name: value.Name,
			StudentId: value.CustomAttrs[0].Value.Text,
			OpenId: value.OpenId,
			Mobile: value.Mobile,
			DepartmentId: value.DepartmentIds[0],
			Avatar: value.Avatar.AvatarOrigin,
		}
		DB.Create(&newStudent)
	}
}

func GetStudentInfo(ctx *gin.Context) {
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

	//定义数据库映射模型 Student
	var student dbModel.Student

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
			"mobile": student.Mobile,
		},
		"msg": "请求成功!",
	})
}

func GetStudentContact(ctx *gin.Context) {
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

	//定义数据库映射模型 Student
	var student dbModel.Student

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
			"mobile": student.Mobile,
		},
		"msg": "请求成功!",
	})
}

