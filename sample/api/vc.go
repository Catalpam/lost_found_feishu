package main

import (
	"context"
	"fmt"
	"lost_found/api/core/request"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/sample/configs"
	vc "lost_found/service/vc/v1"
)

// for redis store and logrus
// configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// configs.TestConfig("https://open.feishu.cn")
var VCService = vc.NewService(configs.TestConfig(constants.DomainFeiShu))

func main() {
	testReserveApply()
}

func testReserveApply() {
	ctx := context.Background()
	coreCtx := core.WrapContext(ctx)
	body := &vc.ReserveApplyReqBody{
		EndTime: 1617161325,
		MeetingSettings: &vc.ReserveMeetingSetting{
			Topic: "Test VC",
			ActionPermissions: []*vc.ReserveActionPermission{{
				Permission: 1,
				PermissionCheckers: []*vc.ReservePermissionChecker{{
					CheckField: 1,
					CheckMode:  1,
					CheckList:  []string{"77bbc392"},
				},
				},
			},
			},
		},
	}
	reqCall := VCService.Reserves.Apply(coreCtx, body, request.SetUserAccessToken("User access token"))
	reqCall.SetUserIdType("user_id")
	result, err := reqCall.Do()
	fmt.Println(coreCtx.GetRequestID())
	fmt.Println(coreCtx.GetHTTPStatusCode())
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return
	}
	fmt.Println(tools.Prettify(result))
}
