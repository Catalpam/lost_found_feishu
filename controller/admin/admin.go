package admin

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetAdmins(ctx *gin.Context)  {
	db := common.GetDB()
	var admins []admin
	var privilages []dbModel.Privilege
	db.Find(&privilages)
	for _,value := range privilages {
		var user dbModel.User
		var permisson string
		db.Where("open_id=?",value.OpenId).First(&user)
		if value.Permission == "1" {
			permisson = common.PlaceId2Name(value.PlaceSmallId)+" "+permissionCode2String(value.Permission)
		} else {
			permisson = permissionCode2String(value.Permission)
		}
		admins = append(admins,admin{
			ID:         value.ID,
			OpenId:     value.OpenId,
			Name:       user.Name,
			Avatar:     user.Avatar,
			Mobile: 	user.Mobile,
			UserId: 	user.StudentId,
			Permission: permisson,
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code" :  200,
		"data": admins,
		"msg": "获取管理员列表成功",
	})
}

type admin struct {
	ID    		uint
	OpenId 		string
	Name		string
	Avatar 		string
	Mobile 		string
	UserId 		string
	Permission 	string
	PlaceName 	string
}