package webController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
)

func DelSubPlace(ctx *gin.Context) {

	key := ctx.Query("key")
	var placeIndex []int
	var place dbModel.Place
	db := common.GetDB()
	var subPlaces []string

	err := json.Unmarshal([]byte(key), &placeIndex)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": err,
			"msg":  "key格式错误！",
		})
		return
	}
	SubIndexInt := placeIndex[2]
	db.Where("campus_id=? AND place_id=?",strconv.Itoa(placeIndex[0]),strconv.Itoa(placeIndex[1])).First(&place)
	json.Unmarshal([]byte(place.Subareas), &subPlaces)
	if len(subPlaces) <= SubIndexInt {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "Key值不合法！",
		})
		return
	}
	willDel := subPlaces[SubIndexInt]
	println("------------删除子地点："+willDel+"------------------")
	subPlaces = append(subPlaces[:SubIndexInt], subPlaces[SubIndexInt+1:]...)
	subareasJson,_ := json.Marshal(subPlaces)

	db.Model(&place).Update("subareas",string(subareasJson))
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除子地点“"+willDel+"”成功",
	})
}
