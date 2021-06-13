package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/comander"
	"lost_found/common"
	"lost_found/dbModel"
	"lost_found/handler"
	"net/http"
	"strings"
	"time"
)

func BulkAddFound(ctx *gin.Context)  {
	db := common.GetDB()

	//获取用户
	var privilage dbModel.Privilege
	var placeSmall dbModel.PlaceSmall
	OpenId := ctx.MustGet("open_id").(string)
	db.Where("open_id=? AND permission=?",OpenId,"1").First(&privilage)
	if privilage.ID == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "你不是地点管理员，如需获取权限，请向管理员申请!  ps：目前应用的管理员有限，管理员可以在本应用的网页管理端配置，具体使用方法请见参赛文档，配置为某一地点管理员后，地点管理员可以使用批量上传。",
		})
		return
	}
	db.Where("id=?",privilage.PlaceSmallId).First(&placeSmall)


	//获取参数
	typeIndex, errType := String2Index(ctx.PostForm("type_index"))

	if errType != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "物品类型参数格式不合法!",
		})
		println("\t\t\t\"code\": 413,\n\t\t\t\"data\": \"\",\n\t\t\t\"msg\":  \"物品类型参数格式不合法!\",\n")
		return
	}

	itemInfo, _ := ctx.GetPostForm("info")
	image, _ := ctx.GetPostForm("image")
	placeDetail,_ := ctx.GetPostForm("place_detail")
	losterInfo, _ := ctx.GetPostForm("loster_info")
	additionalInfo,_ := ctx.GetPostForm("additional_info")

	//检查有无空参数
	{
		if itemInfo == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：info",
			})
			println("\t\t\tctx.JSON(http.StatusOK, gin.H{\n\t\t\t\t\"code\": 413,\n\t\t\t\t\"data\": \"\",\n\t\t\t\t\"msg\":  \"缺少参数：info\",\n\t\t\t})\n")
			return
		}
		if image == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：image",
			})
			println("\t\t\t\t\"code\": 413,\n\t\t\t\t\"data\": \"\",\n\t\t\t\t\"msg\":  \"缺少参数：image\",\n")
			return
		}

		//参数为0，1，2
	}

	//查找index对应的值是否存在于数据库中
	var typeBig   dbModel.TypeBig
	var typeSmall dbModel.TypeSmall
	db.Where("indexx=?", typeIndex[0]).First(&typeBig)
	db.Where("indexx=? AND big_id=?", typeIndex[1], typeBig.ID).First(&typeSmall)

	if typeSmall.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "请求的type_index不存在!",
		})
		println("\t\t\t\"code\": 413,\n\t\t\t\"data\": \"\",\n\t\t\t\"msg\":  \"请求的type_index不存在!\",\n")
		return
	}


	imageNameList := strings.Split(image,`image?name=`)
	imageKey, errUpload := handler.Uploadimage2Feishu(imageNameList[1])
	if errUpload != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": image,
			"msg":  "ImageUrl必须是之前上传过的图片!",
		})
		println(errUpload)
		return
	}
	//将新的Found对象添加至数据库中
	newFound := dbModel.Found{
		Date:          		time.Now().Format("2006-01-02"),
		Time:          		time.Now().Format("15:04"),
		TimeSession:   		Time2Session(),
		OpenId:        		OpenId,
		TypeBigId:          typeBig.ID,
		TypeSmallId:        typeSmall.ID,
		PlaceBigId:         placeSmall.BigId,
		PlaceSmallId:       placeSmall.ID,
		ItemInfo:           itemInfo,
		Image:              image,
		ImageKey: 			imageKey,
		PlaceDetail:        placeDetail,
		CurrentPlace:       "2",
		CurrentPlaceDetail: "失物招领处",
		LosterInfo:         losterInfo,
		AdditionalInfo:     additionalInfo,
	}

	db.Create(&newFound)
	// 新建一个匹配进程，若有符合匹配将会使用机器人发送
	go comander.CheckNewFoundIsMatchedLost(newFound.ID)
	// 此进程继续返回创立Found成功的Response
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": newFound,
		"msg":  "添加Found成功，详见data",
	})
}

func GetPermisson(ctx *gin.Context)  {
	db := common.GetDB()
	var privilage dbModel.Privilege
	OpenId := ctx.MustGet("open_id").(string)
	db.Where("open_id=? AND permission=?",OpenId,"1").First(&privilage)
	if privilage.ID == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "你不是地点管理员，如需获取权限，请向管理员申请!  ps：目前应用的管理员有限，管理员可以在本应用的网页管理端配置，具体使用方法请见参赛文档，配置为某一地点管理员后，地点管理员可以使用批量上传。",
		})
		//ctx.Abort()
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "权限正常，可以使用批量上传！",
		})
	}
}
