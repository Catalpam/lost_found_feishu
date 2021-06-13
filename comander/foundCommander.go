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

	db.Where(&dbModel.Lost{
		TypeSmallId:     	found.TypeSmallId,
		PlaceSmallId1:      found.PlaceSmallId,
		Validity:			true,
	}).Where("match_id=?",0).Find(&losts)
	for _,value := range losts {
		if value.OpenId == found.OpenId {
			continue
		}
		var typeSmall dbModel.TypeSmall
		var placeSmall dbModel.PlaceSmall
		db.Where("id=?",found.TypeSmallId).First(&typeSmall)
		db.Where("id=?",found.PlaceSmallId).First(&placeSmall)
		campus := dbModel.CampusId2Str(placeSmall.CampusId)
		cardMessage.SendCardMessage(
			value.OpenId,
			cardMessage.SuspectedCard(cardMessage.Suspected{
				LostId:      value.ID,
				FoundId:     found.ID,
				ItemSubtype: typeSmall.Name,
				FoundPlace:  campus+" "+placeSmall.BigName+" "+placeSmall.Name,
				FoundDate:   found.Date + " " + common.TimeSessionToChinese(found.TimeSession),
				ImageKey:    found.ImageKey,
			}),
		)
	}

	db.Where(&dbModel.Lost{
		TypeSmallId:     	found.TypeSmallId,
		PlaceSmallId2:      found.PlaceSmallId,
		Validity:			true,
	}).Where("match_id=?",0).Find(&losts)
	for _,value := range losts {
		if value.OpenId == found.OpenId {
			continue
		}
		var typeSmall dbModel.TypeSmall
		var placeSmall dbModel.PlaceSmall
		db.Where("id=?",found.TypeSmallId).First(&typeSmall)
		db.Where("id=?",found.PlaceSmallId).First(&placeSmall)
		campus := ""

		cardMessage.SendCardMessage(
			value.OpenId,
			cardMessage.SuspectedCard(cardMessage.Suspected{
				LostId:      value.ID,
				FoundId:     found.ID,
				ItemSubtype: typeSmall.Name,
				FoundPlace:  campus+" "+placeSmall.BigName+" "+placeSmall.Name,
				FoundDate:   found.Date + " " + common.TimeSessionToChinese(found.TimeSession),
				ImageKey:    found.ImageKey,
			}),
		)
	}

	db.Where(&dbModel.Lost{
		TypeSmallId:     	found.TypeSmallId,
		PlaceSmallId3:      found.PlaceSmallId,
		Validity:			true,
	}).Where("match_id=?",0).Find(&losts)
	for _,value := range losts {
		if value.OpenId == found.OpenId {
			continue
		}
		var typeSmall dbModel.TypeSmall
		var placeSmall dbModel.PlaceSmall
		db.Where("id=?",found.TypeSmallId).First(&typeSmall)
		db.Where("id=?",found.PlaceSmallId).First(&placeSmall)
		campus := ""

		cardMessage.SendCardMessage(
			value.OpenId,
			cardMessage.SuspectedCard(cardMessage.Suspected{
				LostId:      value.ID,
				FoundId:     found.ID,
				ItemSubtype: typeSmall.Name,
				FoundPlace:  campus+" "+placeSmall.BigName+" "+placeSmall.Name,
				FoundDate:   found.Date + " " + common.TimeSessionToChinese(found.TimeSession),
				ImageKey:    found.ImageKey,
			}),
		)
	}
	return
}

func SendThx(MatchId uint)  {
	db := common.GetDB()
	var match dbModel.Match
	db.Where("id=?",MatchId).First(&match)

	cardMessage.SendCardMessage(
		match.FoundOpenId,
		cardMessage.FoundClaimCard(cardMessage.FoundClaim{
			ItemSubtype:  match.TypeName,
			LeaveMessage: match.LosterComment,
			ImageKey:     match.ImageKey,
		}),
	)
	println("感谢信息发送成功！")

	cardMessage.SendCardMessage(
		match.LosterOpenId,
		cardMessage.ThanksHasSendCard(cardMessage.ThanksHasSend{
			ItemSubtype: match.TypeName,
			FoundDate:   match.FoundDate,
			ImageKey:    match.ImageKey,
		}),
	)
	println("失主反馈信息发送成功！")
}

func SendUesrToBoth(MatchId uint)  {
	db := common.GetDB()
	var match dbModel.Match
	var user  dbModel.User
	db.Where("id=?",MatchId).First(&match)

	// 若不为自己带走，Return
	if match.CurrentPlace != "1" {
		return
	}
	db.Where("open_id=?",match.FoundOpenId).First(&user)
	cardMessage.SendCardMessage(
		match.LosterOpenId,
		cardMessage.SendUser2LosterCard(cardMessage.SendUser2Loster{
			FounderName: user.Name,
			ItemSubtype: match.TypeName,
			FoundDate:   match.FoundDate,
			ImageKey:    match.ImageKey,
		}),
	)
	println("---------------联系方式发送给Loster成功！-------------------")

	cardMessage.SendCardMessage(
		match.FoundOpenId,
		cardMessage.SendUser2FounderCard(cardMessage.SendUser2Founder{
			ItemSubtype: match.TypeName,
			FoundDate:   match.FoundDate,
			ImageKey:    match.ImageKey,
		}),
	)
	println("---------------联系方式发送给Founder成功！---------------")
	cardMessage.ImSendUser(match.FoundOpenId,match.LosterOpenId)
	cardMessage.ImSendUser(match.LosterOpenId,match.FoundOpenId)
}

func SelectLosterName(found dbModel.Found) {
	fmt.Printf("--------开始通过姓名查找。姓名为：%s----------",found.LosterInfo)
	db := common.GetDB()
	var users []dbModel.User
	db.Where("name=?",found.LosterInfo).Find(&users)
	if len(users) == 0 {
		fmt.Printf("--------姓名为%s的用户不存在----------",found.LosterInfo)
		return
	}
	println(found.ID)
	println(found.LosterInfo)
	var placeSmall dbModel.PlaceSmall
	db.Where("id=?",found.PlaceSmallId).First(&placeSmall)
	campus := dbModel.CampusId2Str(placeSmall.CampusId)
	println(placeSmall.Name)
	println(placeSmall.BigName)

	for _, value := range users {
		cardMessage.SendCardMessage(
			value.OpenId,
			cardMessage.SameNameCard(cardMessage.SameName{
				FoundId:    found.ID,
				FoundPlace: campus+" "+placeSmall.BigName +" "+placeSmall.Name,
				ImageKey:   found.ImageKey,
			}),
		)
	}
}