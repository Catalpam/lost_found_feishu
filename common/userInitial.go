package common

import (
	"context"
	"fmt"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/dbModel"
	contact "lost_found/service/contact/v3"
)

var contactService = contact.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func UpdateStudentList() {
	DB := GetDB()
	coreCtx := core.WrapContext(context.Background())
	reqCall := contactService.Users.List(coreCtx)
	result, err := reqCall.Do()
	if err != nil {
		e := err.(*response.Error)
		fmt.Printf(tools.Prettify(e))
		return
	}
	for _, value := range result.Items {
		var dbStudent dbModel.User
		open_id := value.OpenId
		var id string
		if len(value.CustomAttrs) != 0 {
			// id = ""
			id = value.CustomAttrs[0].Value.Text
		} else {
			id = "非电子科大成员，无学号/工号"
		}
		DB.Where("open_id = ?", open_id).First(&dbStudent)
		//更新学生信息：
		if dbStudent.ID != 0 {
			if dbStudent.Name != value.Name {
				dbStudent.Name = value.Name
			}
			if dbStudent.StudentId != id {
				dbStudent.StudentId = id
			}
			if dbStudent.Mobile != value.Mobile {
				dbStudent.Mobile = value.Mobile
			}
			if dbStudent.DepartmentId != value.DepartmentIds[0] {
				dbStudent.DepartmentId = value.DepartmentIds[0]
			}
			if dbStudent.Avatar != value.Avatar.AvatarOrigin {
				dbStudent.Avatar = value.Avatar.AvatarOrigin
			}
			//新建完成后Continue
			continue
		}
		//创建新学生
		newStudent := dbModel.User{
			Name:         value.Name,
			StudentId:    id,
			OpenId:       value.OpenId,
			Mobile:       value.Mobile,
			DepartmentId: value.DepartmentIds[0],
			Avatar:       value.Avatar.AvatarOrigin,
		}
		DB.Create(&newStudent)
	}
}
