package comander

import (
	"fmt"
	"lost_found/common"
	"lost_found/dbModel"
	"strconv"
)

const FoundedBySelf uint64 = 4294967294

func HasFounded(lostId uint64) bool{
	db := common.GetDB()
	var lost dbModel.Lost
	fmt.Printf("-------开始将%s的状态变为自己找到---------\n", strconv.FormatUint(lostId, 10))
	db.Where("id=? AND match_id=?",lostId,0).First(&lost)
	if lost.ID == 0 {
		fmt.Printf("-------未匹配的%s不存在---------\n", strconv.FormatUint(lostId, 10))
		return false
	} else {
		db.Update("match_id", FoundedBySelf)
		fmt.Printf("-------成功将%s的状态变为自己找到---------\n", strconv.FormatUint(lostId, 10))
		fmt.Printf("-------成功将%s的状态变为自己找到---------\n", strconv.FormatUint(lostId, 10))
		return true
	}
}