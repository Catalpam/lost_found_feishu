package comander

import (
	"fmt"
	"lost_found/cardMessage"
	"lost_found/common"
	"lost_found/dbModel"
	"strconv"
)

const FoundedBySelf uint64 = 4294967294

func HasFounded(lostId uint64) bool{
	db := common.GetDB()
	var lost dbModel.Lost
	var message = ""
	fmt.Printf("-------开始将%s的状态变为自己找到---------\n", strconv.FormatUint(lostId, 10))
	db.Where("id=? AND match_id=?",lostId,0).First(&lost)
	if lost.ID == 0 {
		fmt.Printf("-------未匹配的%s不存在---------\n", strconv.FormatUint(lostId, 10))
		db.Where("id=?",lostId).First(&lost)
		message = fmt.Sprintf("失物【%s】已经登记为已找到啦！请勿重复操作。如有疑问，请联系管理员。",lost.TypeSubName)
		cardMessage.SendMessage(lost.LosterOpenId,message)
		return false
	} else {
		db.Model(&lost).Update("match_id", FoundedBySelf)
		fmt.Printf("-------成功将%s的状态变为自己找到---------\n", strconv.FormatUint(lostId, 10))
		fmt.Printf("-------成功将%s的状态变为自己找到---------\n", strconv.FormatUint(lostId, 10))
		message = fmt.Sprintf("您的失物【%s】状态已更改为已找到。恭喜您找到了您的失物",lost.TypeSubName)
		cardMessage.SendMessage(lost.LosterOpenId,message)
		return true
	}
}