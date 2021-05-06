package main

import (
	"context"
	"fmt"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/sample/configs"
	contact "lost_found/service/contact/v3"
)

// for redis store and logrus
// configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// configs.TestConfig("https://open.feishu.cn")
var contactService = contact.NewService(configs.TestConfig(constants.DomainFeiShu))

func main() {
	testUserServiceList()
	//testDepartmentServiceList()
}
func testUserServiceList() {
	coreCtx := core.WrapContext(context.Background())
	reqCall := contactService.Users.List(coreCtx)
	reqCall.SetDepartmentIdType("open_id")
	reqCall.SetPageSize(20)
	reqCall.SetDepartmentIdType("open_department_id")
	reqCall.SetDepartmentId("0")
	reqCall.SetUserIdType("open_id")
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

func testDepartmentServiceList() {
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
