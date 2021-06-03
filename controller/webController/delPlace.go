package webController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
)

func DelPlace(ctx *gin.Context) {

	key := ctx.Query("key")
	var placeIndex []int
	var place dbModel.Place
	db := common.GetDB()

	err := json.Unmarshal([]byte(key), &placeIndex)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": err,
			"msg":  "key格式错误！",
		})
		return
	}
	db.Where("campus_id=? AND place_id=?",strconv.Itoa(placeIndex[0]),strconv.Itoa(placeIndex[1])).First(&place)
	willDel := place.Name
	db.Delete(&place)
	println("------------删除父地点："+willDel+"------------------")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除父地点“"+willDel+"”成功",
	})
}
