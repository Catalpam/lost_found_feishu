package comander

import (
	"fmt"
	"lost_found/cardMessage"
	"lost_found/common"
	"lost_found/dbModel"
)

func CheckNewFoundIsMatchedLost (newFoundId uint) {
	db := common.GetDB()
	var found dbModel.Found
	var losts []dbModel.Lost
	fmt.Printf("开始匹配，FoundId为：%d\n",newFoundId)
	db.Where("id=?",newFoundId).First(&found)
	if found.ID == 0 {
		panic("Error：新的Found没有找到！")
	}
	placeIndexStr:= "[\""+found.Place+"\",\""+found.SubPlace+"\"]"
	fmt.Printf("匹配的地点为：%s\n",placeIndexStr)
	db.Where(&dbModel.Lost{
		TypeSubName:     found.SubType,
		LostPlace1:      placeIndexStr ,
		//LostDate:        found.FoundDate,
		//LostTimeSession: found.FoundTimeSession,
	}).Where("match_id=?",0).Find(&losts)
	for _,value := range losts {
		if value.LosterOpenId == found.FoundOpenId {
			continue
		}
		cardMessage.SendCardMessage(
			value.LosterOpenId,
			cardMessage.SuspectedCard(cardMessage.Suspected{
				LostId:      value.ID,
				FoundId:     found.ID,
				ItemSubtype: found.SubType,
				FoundPlace:  found.Campus+" "+found.SubPlace,
				FoundDate:   found.FoundDate + " " + common.TimeSessionToChinese(found.FoundTimeSession),
				ImageKey:    found.ImageKey,
			}),
		)
	}
	db.Where(&dbModel.Lost{
		TypeSubName:     found.SubType,
		LostPlace3:      placeIndexStr,
		LostDate:        found.FoundDate,
		LostTimeSession: found.FoundTimeSession,
	}).Where("match_id=?",0).Find(&losts)
	for _,value := range losts {
		if value.LosterOpenId == found.FoundOpenId {
			continue
		}
		cardMessage.SendCardMessage(
			value.LosterOpenId,
			cardMessage.SuspectedCard(cardMessage.Suspected{
				LostId:      value.ID,
				FoundId:     found.ID,
				ItemSubtype: found.SubType,
				FoundPlace:  found.Campus+" "+found.SubPlace,
				FoundDate:   found.FoundDate + " " + common.TimeSessionToChinese(found.FoundTimeSession),
				ImageKey:    found.ImageKey,
			}),
		)
	}
	db.Where(&dbModel.Lost{
		TypeSubName:     found.SubType,
		LostPlace2:      placeIndexStr,
		LostDate:        found.FoundDate,
		LostTimeSession: found.FoundTimeSession,
	}).Where("match_id=?",0).Find(&losts)
	for _,value := range losts {
		if value.LosterOpenId == found.FoundOpenId {
			continue
		}
		cardMessage.SendCardMessage(
			value.LosterOpenId,
			cardMessage.SuspectedCard(cardMessage.Suspected{
				LostId:      value.ID,
				FoundId:     found.ID,
				ItemSubtype: found.SubType,
				FoundPlace:  found.Campus+" "+found.SubPlace,
				FoundDate:   found.FoundDate + " " + common.TimeSessionToChinese(found.FoundTimeSession),
				ImageKey:    found.ImageKey,
			}),
		)
	}
	return
}
