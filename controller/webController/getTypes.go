package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetTypesSmall(ctx *gin.Context) {
	db := common.GetDB()
	var typeSmalls []dbModel.TypeSmall
	var retTypes []RetSmallType
	db.Order("big_id ASC").Order("indexx ASC").Find(&typeSmalls)

	for _, value := range typeSmalls {
		retTypes = append(retTypes, RetSmallType{
			BigName:    value.BigName,
			SmallName: 	value.Name,
			Key:      	value.ID,
			BigKey:	  	value.BigId,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": retTypes,
		"msg":  "Place获取成功",
	})
}

func GetTypesBig(ctx *gin.Context) {
	db := common.GetDB()
	var bigs []dbModel.TypeBig
	var options []RetBigType
	db.Order("indexx ASC").Find(&bigs)

	for _, value := range bigs {
		options = append(options,RetBigType{
			Type:   value.Name,
			BigKey: value.ID,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": options,
		"msg":  "BigTypes获取成功",
	})
}


type RetSmallType struct {
	BigName 		string
	SmallName 		string
	Key        		uint
	BigKey			uint
}

type RetBigType struct {
	Type 		string `json:"label"`
	BigKey      uint   `json:"value"`
}