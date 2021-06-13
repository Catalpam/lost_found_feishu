package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
	"strings"
)

func GetFoundList(ctx *gin.Context) {
	db := common.GetDB()
	var founds []dbModel.Found
	// 获取Form中的参数 FoundId
	FoundIdStr := ctx.PostForm("id")
	if FoundIdStr == "" {
		FoundIdStr = ctx.PostForm("found_id")
	}
	// 查找参数
	if FoundIdStr != "" {
		FoundId, err := strconv.ParseUint(ctx.PostForm("id"), 10 ,32)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": err,
				"msg":  "id格式不合法！",
			})
		}
		println(FoundId)
		db.Where("id=?",FoundId).Find(&founds)
	} else {
		if ctx.PostForm("type_index") == "" && ctx.PostForm("place_index") == ""{
			db.Where("match_id=?",0).Find(&founds)
		} else {
			SelectFound(&founds,ctx)
		}
	}
	returnFounds(&founds,ctx)
}


func returnFounds(founds *[]dbModel.Found, ctx *gin.Context)  {
	if len(*founds) == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data":"[]",
			"msg":  "Ops,没有查询到符合条件的Founds",
		})
		return
	}
	var db = common.GetDB()
	var FoundList []FoundListModel
	for _, value := range *founds {
		var placeSmall dbModel.PlaceSmall
		var typeSmall  dbModel.TypeSmall
		db.Where("id=?",value.PlaceSmallId).First(&placeSmall)
		db.Where("id=?",value.TypeSmallId).First(&typeSmall)

		tempFound := FoundListModel{
			ID:        			value.ID,
			SubType:   			typeSmall.BigName+" "+typeSmall.Name,
			Campus:    			dbModel.CampusId2Str(placeSmall.CampusId),
			Place:     			placeSmall.BigName+"-"+placeSmall.Name,
			PlaceDetail: 		value.PlaceDetail,
			Image:	   			value.Image,
			FoundDate: 		 	value.Date,
			FoundTime: 		 	value.Time,
			Info: 		 		value.ItemInfo,
			AdditionalInfo : 	value.AdditionalInfo,
		}
		FoundList = append(FoundList, tempFound)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": FoundList,
		"msg":  "获取Found List成功",
	})
}

type FoundListModel struct {
	ID uint
	SubType string
	// Location
	Campus string
	Place string
	PlaceDetail string
	// Image
	Image string
	ImageList string
	// Time
	FoundDate string
	FoundTime string
	Info string
	AdditionalInfo string
}

func SelectFound(founds *[]dbModel.Found, ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	campusId, _ := ctx.GetPostForm("campus_id")
	//date, _ := ctx.GetPostForm("date")
	typeIndex := [2]uint{0, 0}
	placeIndex := [2]uint{0, 0}
	var errType, errPlace error
	typeIndex, errType = String2Index(ctx.PostForm("type_index"))
	placeIndex, errPlace = String2Index(ctx.PostForm("place"))
	println(placeIndex[1])
	//查找index对应的值是否存在于数据库中
	var typeBig   = dbModel.TypeBig{ID: 0}
	var typeSmall = dbModel.TypeSmall{ID: 0}
	if errType == nil {
		db.Where("indexx=?", typeIndex[0]).First(&typeBig)
		db.Where("indexx=? AND big_id=?", typeIndex[1], typeBig.ID).First(&typeSmall)
	}

	var placeBig   = dbModel.PlaceBig{ID: 0}
	var placeSmall = dbModel.PlaceSmall{ID: 0}
	if errPlace == nil && campusId != ""{
		db.Where("indexx=? AND campus_id=?", placeIndex[0], campusId).First(&placeBig)
		db.Where("indexx=? AND big_id=?", placeIndex[1], typeBig.ID).First(&placeSmall)
	} else if errPlace == nil && campusId == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "啊哦，没有填写校区哦!",
		})
	} else if errPlace != nil && campusId != "" {
		db.Where("campus_id=?", campusId).First(&placeBig)
		db.Where("indexx=? AND big_id=?", placeIndex[1], typeBig.ID).First(&placeSmall)
	}

	println(placeSmall.Indexx)
	println(placeSmall.Name)

	//在数据库中查找Found对象
	println("----------------准备开始查找符合条件的Found----------------------")
	db.Where(&dbModel.Found{
		TypeSmallId:   typeSmall.ID,
		PlaceSmallId:  placeSmall.ID,
		//Date:          date,
	}).Where("match_id=?",0).Find(&founds)
}

func String2Index(str string) ([2]uint,error){
	countSplit := strings.Split(str, ",")
	var index [2]uint
	var index64 [2]uint64
	var err error
	index64[0],err = strconv.ParseUint(countSplit[0], 10, 32)
	if err != nil {
		println(err)
		return [2]uint{0,0}, err
	}
	index64[1],err = strconv.ParseUint(countSplit[1], 10, 32)
	if err != nil {
		println(err)
		return [2]uint{0,0}, err
	}
	index[0] = uint(index64[0])
	index[1] = uint(index64[1])
	return index, nil
}