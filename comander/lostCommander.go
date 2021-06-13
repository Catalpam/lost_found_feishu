package comander

import (
	"fmt"
	"lost_found/cardMessage"
	"lost_found/common"
	"lost_found/dbModel"
	"strconv"
	"time"
)

func HasFounded(lostId uint64) bool{
	db := common.GetDB()
	var lost 		dbModel.Lost
	var typeSmall 	dbModel.TypeSmall
	var message = ""
	fmt.Printf("-------开始将%s的状态变为自己找到---------\n", strconv.FormatUint(lostId, 10))
	db.Where("id=? AND match_id=?",lostId,0).First(&lost)
	db.Where("id=?",lost.TypeSmallId).First(&typeSmall)
	if lost.ID == 0 {
		fmt.Printf("-------未匹配的%s不存在---------\n", strconv.FormatUint(lostId, 10))
		db.Where("id=?",lostId).First(&lost)
		message = fmt.Sprintf("失物【%s】已经登记为已找到啦！请勿重复操作。如有疑问，请联系管理员。",typeSmall.Name)
		cardMessage.SendMessage(lost.OpenId,message)
		return false
	} else {
		newMatch := dbModel.Match{
			FoundDate:          time.Now().Format("2006-01-02"),
			Time:          		time.Now().Format("15:04"),
			TimeSession:   		Time2Session(),
			LosterOpenId:		lost.OpenId,
			FoundOpenId:  		lost.OpenId,
			TypeBigId:   		lost.TypeBigId,
			TypeSmallId:  		lost.TypeSmallId,
			TypeName:     		common.TypeId2Name(lost.TypeSmallId),
			PlaceName:    		"自行找到",
			Image:              "https://",
		}
		db.Create(&newMatch)
		db.Model(&lost).Update("is_found_by_self", true)
		db.Model(&dbModel.Match{

		})
		fmt.Printf("-------成功将%s的状态变为自己找到---------\n", strconv.FormatUint(lostId, 10))
		fmt.Printf("-------成功将%s的状态变为自己找到---------\n", strconv.FormatUint(lostId, 10))
		message = fmt.Sprintf("您的失物【%s】状态已更改为已找到。恭喜您找到了您的失物",typeSmall.Name)
		cardMessage.SendMessage(lost.OpenId,message)
		return true
	}
}

func Time2Session() string  {
	currentTime := time.Now().Hour()
	var currentSession string

	switch currentTime {
	case 6,7,8,9,10:
		currentSession = "morning"
	case 11,12,13:
		currentSession = "noon"
	case 14,15,16,17,18:
		currentSession = "afternoon"
	case 19,20,21:
		currentSession = "evening"
	case 22,23,1,2,3,4,5:
		currentSession = "night"
	}
	return currentSession
}