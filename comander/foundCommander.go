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
	if found.LosterInfo != "" {
		go SelectLosterName(found)
	}
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

func SendThx(FoundId uint)  {
	db := common.GetDB()
	var found dbModel.Found
	var lost  dbModel.Lost
	db.Where("id=?",FoundId).First(&found)
	db.Where("id=?",found.MatchId).First(&lost)
	if lost.ID == 0 || found.ID == 0 {
		println("不存在对应Found或Lost")
		return
	}

	cardMessage.SendCardMessage(
		found.FoundOpenId,
		cardMessage.FoundClaimCard(cardMessage.FoundClaim{
			ItemSubtype:  found.SubType,
			LeaveMessage: found.LosterComment,
			ImageKey:     found.ImageKey,
		}),
	)
	println("感谢信息发送成功！")

	cardMessage.SendCardMessage(
		lost.LosterOpenId,
		cardMessage.ThanksHasSendCard(cardMessage.ThanksHasSend{
			ItemSubtype: found.SubType,
			FoundDate:   found.FoundDate,
			ImageKey:    found.ImageKey,
		}),
	)
	println("失主反馈信息发送成功！")
}

func SendUesrToBoth(FoundId uint)  {
	db := common.GetDB()
	var found dbModel.Found
	var lost  dbModel.Lost
	var user  dbModel.User
	db.Where("id=?",FoundId).First(&found)
	db.Where("id=?",found.MatchId).First(&lost)
	if lost.ID == 0 || found.ID == 0 {
		println("不存在对应Found或Lost")
		return
	}
	// 若不为自己带走，Return
	if found.CurrentPlace != "1" {
		return
	}
	db.Where("open_id=?",found.FoundOpenId).First(&user)
	cardMessage.SendCardMessage(
		lost.LosterOpenId,
		cardMessage.SendUser2LosterCard(cardMessage.SendUser2Loster{
			FounderName: user.Name,
			ItemSubtype: found.SubType,
			FoundDate:   found.FoundDate,
			ImageKey:    found.ImageKey,
		}),
	)
	println("---------------联系方式发送给Loster成功！-------------------")

	cardMessage.SendCardMessage(
		found.FoundOpenId,
		cardMessage.SendUser2FounderCard(cardMessage.SendUser2Founder{
			ItemSubtype: found.SubType,
			FoundDate:   found.FoundDate,
			ImageKey:    found.ImageKey,
		}),
	)
	println("---------------联系方式发送给Founder成功！---------------")
	cardMessage.ImSendUser(found.FoundOpenId,lost.LosterOpenId)
	cardMessage.ImSendUser(lost.LosterOpenId,found.FoundOpenId)
}

func SelectLosterName(found dbModel.Found) {
	fmt.Printf("--------开始通过姓名查找。姓名为：%s----------",found.LosterInfo)
	db := common.GetDB()
	var user dbModel.User
	db.Where("name=?",found.LosterInfo).First(&user)
	if user.ID == 0 {
		fmt.Printf("--------姓名为%s的用户不存在----------",found.LosterInfo)
		return
	}
	cardMessage.SendCardMessage(
		user.OpenId,
		cardMessage.SameNameCard(cardMessage.SameName{
			FoundId:    found.ID,
			FoundPlace: found.SubPlace,
			ImageKey:   found.ImageKey,
		}),
	)
}