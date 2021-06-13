package admin

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func SetAdminPrivilage(ctx *gin.Context) {
	OpenId := ctx.MustGet("open_id").(string)
	newOpenId := ctx.Query("open_id")
	newAdminCode := ctx.Query("privilage")
	db := common.GetDB()
	if !(newAdminCode == "2" || newAdminCode == "3")  {
		ctx.JSON(http.StatusOK,gin.H{
			"code":403,
			"msg":"权限类型不合法！",
		})
		return
	}
	var privilage dbModel.Privilege
	db.Where("open_id=?",OpenId).First(&privilage)
	if privilage.Permission != "3" {
		ctx.JSON(http.StatusOK,gin.H{
			"code":413,
			"msg":"权限不足，只有超级管理员才能更改其他用户的权限！",
		})
		return
	}
	var newPrivilage dbModel.Privilege
	db.Where("open_id=?",newOpenId).First(&newPrivilage)

	if newPrivilage.ID == 0{
		db.Create(&dbModel.Privilege{
			OpenId:     newOpenId,
			Permission: newAdminCode,
		})
	} else {
		if newPrivilage.Permission != newAdminCode {
			db.Model(&newPrivilage).Update("permission",newAdminCode)
		} else {
			ctx.JSON(http.StatusOK,gin.H{
				"code":413,
				"msg":"该用户已经是"+permissionCode2String(newAdminCode)+"！",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"成功将该用户的权限更改为"+permissionCode2String(newAdminCode)+"！",
	})
	return
}

func SetPlaceAdmin(ctx *gin.Context) {
	OpenId := ctx.MustGet("open_id").(string)
	newOpenId := ctx.Query("open_id")
	placeSmallId := ctx.Query("place_key")
	db := common.GetDB()
	var privilage dbModel.Privilege
	db.Where("open_id=?",OpenId).First(&privilage)
	if privilage.Permission != "3" {
		ctx.JSON(http.StatusOK,gin.H{
			"code":413,
			"msg":"权限不足，只有超级管理员才能更改其他用户的权限！",
		})
		return
	}

	var newPrivilage dbModel.Privilege
	db.Where("open_id=?",newOpenId).First(&newPrivilage)
	placeName := ""
	var placeSmall dbModel.PlaceSmall
	db.Where("id=?",placeSmallId).Find(&placeSmall)
	if placeSmall.ID == 0 {
		ctx.JSON(http.StatusOK,gin.H{
			"code":403,
			"msg":"选择的地点不存在！",
		})
		return
	} else {
		placeName =  placeSmall.BigName+"-"+placeSmall.Name
	}

	if newPrivilage.ID == 0{
		db.Create(&dbModel.Privilege{
			OpenId:     newOpenId,
			Permission: "1",
			PlaceSmallId: placeSmall.ID,
		})
	} else {
		if newPrivilage.Permission != "1"  ||  newPrivilage.PlaceSmallId != placeSmall.ID {
			db.Model(&newPrivilage).Update("permission","1").Update("place_small_id",placeSmallId)
		} else {
			ctx.JSON(http.StatusOK,gin.H{
				"code":413,
				"msg":"该用户已经是“"+placeName+"”的地点管理员！",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"成功将该用户的权限更改为“"+placeName+"”的地点管理员！",
	})
	return
}




func GetPrivilages(ctx *gin.Context) {
	var options []PrivilageType

	options = append(options,PrivilageType{
		Privilage: "超级管理员",
		BigKey:    "3",
	})
	options = append(options,PrivilageType{
		Privilage: "平台审核管理员",
		BigKey:    "2",
	})
	options = append(options,PrivilageType{
		Privilage: "地点管理员",
		BigKey:    "1",
	})
	options = append(options,PrivilageType{
		Privilage: "普通用户",
		BigKey:    "0",
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": options,
		"msg":  "管理员权限类型获取成功",
	})
}

type PrivilageType struct {
	Privilage 	string 	`json:"label"`
	BigKey      string  `json:"value"`
	Info      	string 	`json:"info"`
}

func SearchUser(ctx *gin.Context)  {
	db := common.GetDB()
	searchKey := ctx.Query("key")
	var queries = [4]string{"open_id","student_id","name","mobile"}
	var retUsers []retUser
	var users []dbModel.User

	println("正在以"+queries[0]+"作为关键词检索")
	db.Where("open_id=?",searchKey).Find(&users)
	for _,value := range users {
		retUsers = append(retUsers,retUser{
			Name:      value.Name,
			UserId:    value.StudentId,
			OpenId:    value.OpenId,
			Mobile:    value.Mobile,
			Avatar:    value.Avatar,
		})
	}

	println("正在以"+queries[1]+"作为关键词检索")
	db.Where("student_id=?",searchKey).Find(&users)
	for _,value := range users {
		retUsers = append(retUsers,retUser{
			Name:      value.Name,
			UserId:    value.StudentId,
			OpenId:    value.OpenId,
			Mobile:    value.Mobile,
			Avatar:    value.Avatar,
		})
	}

	println("正在以"+queries[2]+"作为关键词检索")
	db.Where("name=?",searchKey).Find(&users)
	for _,value := range users {
		retUsers = append(retUsers,retUser{
			Name:      value.Name,
			UserId:    value.StudentId,
			OpenId:    value.OpenId,
			Mobile:    value.Mobile,
			Avatar:    value.Avatar,
		})
	}

	println("正在以"+queries[3]+"作为关键词检索")
	db.Where("mobile=?",searchKey).Find(&users)
	for _,value := range users {
		retUsers = append(retUsers,retUser{
			Name:      value.Name,
			UserId:    value.StudentId,
			OpenId:    value.OpenId,
			Mobile:    value.Mobile,
			Avatar:    value.Avatar,
		})
	}

	if len(retUsers) == 0 {
		ctx.JSON(http.StatusOK,gin.H{
			"code":403,
			"data":nil,
			"msg":"没有找到对应的用户，请检查拼写，如如姓名拼写是否正确",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":retUsers,
		"msg":"搜索成功",
	})
}

type retUser struct {
	Name string
	UserId string
	OpenId string
	Mobile string
	Avatar string
}

func permissionCode2String(code string)string {
	if code == "0" {return "普通用户"}
	if code == "1" {return "地点管理员"}
	if code == "2" {return "平台审核管理员"}
	if code == "3" {return "超级管理员"}
	return ""
}

