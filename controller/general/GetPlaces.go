package general

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetPlaces(ctx *gin.Context) {
	campusId := ctx.PostForm("campus_id")
	db := common.GetDB()
	var placeBig []dbModel.PlaceBig

	db.Where("campus_id = ?", campusId).Order("indexx asc").Find(&placeBig)

	var placeBigs []string
	var placeSmalls []placeSmallIndex
	for _, value := range placeBig {
		var placeSmall 	[]dbModel.PlaceSmall
		var placeSmallSingle placeSmallIndex
		db.Where("big_id=?",value.ID).Order("indexx asc").Find(&placeSmall)
		for _, subValue := range placeSmall {

			placeSmallSingle = append(placeSmallSingle,subValue.Name)
		}
		if placeSmallSingle != nil {
			placeBigs = append(placeBigs, value.Name)
			placeSmalls = append(placeSmalls,placeSmallSingle)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"place1": placeBigs,
			"place2": placeSmalls,
		},
		"msg":  "Places返回成功",
	})
}

type placeSmallIndex []string

func GetTypes(ctx *gin.Context)  {
	db := common.GetDB()
	var typeBig	 	[]dbModel.TypeBig
	var typeBigs 	[]string

	db.Order("indexx").Find(&typeBig)

	var typeSmalls []typeSmallIndex
	for _, value := range typeBig {
		var typeSmall 	[]dbModel.TypeSmall
		var smallSingle typeSmallIndex
		db.Where("big_id=?",value.ID).Order("indexx").Find(&typeSmall)
		for _, subValue := range typeSmall {
			smallSingle = append(smallSingle,subValue.Name)
		}
		if smallSingle != nil {
			typeBigs = append(typeBigs, value.Name)
			typeSmalls = append(typeSmalls,smallSingle)
		}
	}

	ctx.JSON(http.StatusOK,gin.H{
		"code": 200,
		"data": gin.H{
			"type1":typeBigs,
			"type2":typeSmalls,
		},
		"msg": "物品类型Types返回成功",
	})
}

type typeSmallIndex []string

func PlaceId2Name(id uint) string  {
	db := common.GetDB()
	var placeSmall dbModel.PlaceSmall
	println(id)
	db.Where("id=?",id).First(&placeSmall)
	if placeSmall.ID == 0 {
		return ""
	} else {
		return placeSmall.BigName+" "+placeSmall.Name
	}
}