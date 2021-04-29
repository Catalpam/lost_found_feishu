package feishu

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/api/core/response"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/constants"
	"github.com/larksuite/oapi-sdk-go/core/tools"
	contact "github.com/larksuite/oapi-sdk-go/service/contact/v3"
	"lost_found/common"
	"lost_found/dbModel"
	"lost_found/feishuApi/configs"
)

// for redis store and logrus
// configs.ConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// configs.FeishuConfig("https://open.feishu.cn")
var contactService = contact.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func UpdateStudentList() {
	DB := common.GetDB()
	coreCtx := core.WrapContext(context.Background())
	reqCall := contactService.Users.List(coreCtx)
	result, err := reqCall.Do()
	if err != nil {
		e := err.(*response.Error)
		fmt.Printf(tools.Prettify(e))
		return
	}
	print("\n\n\n\n\n\n\n\n\n\\n\n\n\\n\n")

	for _, value := range result.Items {
		newStudent := dbModel.Student{
			Name: value.Name,
			StudentId: value.CustomAttrs[0].Value.Text,
			OpenId: value.OpenId,
			Mobile: value.Mobile,
			DepartmentId: value.DepartmentIds[0],
			Avatar: value.Avatar.AvatarOrigin,
		}
		DB.Create(&newStudent)
	}
}

func UserServiceList() {
	coreCtx := core.WrapContext(context.Background())
	reqCall := contactService.Users.List(coreCtx)
	//reqCall.SetDepartmentIdType("open_id")
	//reqCall.SetPageSize(20)
	//reqCall.SetDepartmentIdType("open_department_id")
	//reqCall.SetDepartmentId("0")
	//reqCall.SetUserIdType("open_id")
	result, err := reqCall.Do()
	fmt.Printf("request_id:%s", coreCtx.GetRequestID())
	fmt.Printf("http status code:%d", coreCtx.GetHTTPStatusCode())
	if err != nil {
		e := err.(*response.Error)
		fmt.Printf(tools.Prettify(e))
		return
	}
	fmt.Printf("reault:%s", tools.Prettify(result))
}


func DepartmentServiceList() {
	coreCtx := core.WrapContext(context.Background())
	reqCall := contactService.Departments.List(coreCtx)
	reqCall.SetDepartmentIdType("open_department_id")
	reqCall.SetUserIdType("open_id")
	result, err := reqCall.Do()
	fmt.Printf("request_id:%s\n", coreCtx.GetRequestID())
	fmt.Printf("http status code:%d", coreCtx.GetHTTPStatusCode())
	if err != nil {
		e := err.(*response.Error)
		fmt.Printf(tools.Prettify(e))
		return
	}
	fmt.Printf("reault:%s", tools.Prettify(result))
}
