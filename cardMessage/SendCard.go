package cardMessage

import (
	"context"
	"encoding/json"
	"fmt"
	"lost_found/api"
	"lost_found/api/core/request"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/tools"
)

func SendCardMessage(OpenID string, cardContent string) {
	coreCtx := core.WrapContext(context.Background())
	card := map[string]interface{}{}
	println(cardContent)
	err := json.Unmarshal([]byte(cardContent), &card)
	if err != nil {
		panic(err)
	}
	body := map[string]interface{}{
		"open_id":  OpenID,
		"msg_type": "interactive",
		"card":     card,
	}
	ret := make(map[string]interface{})
	req := request.NewRequestWithNative("message/v4/send", "POST",
		request.AccessTokenTypeTenant, body, &ret,
	)
	err = api.Send(coreCtx, conf, req)
	fmt.Println(coreCtx.GetRequestID())
	fmt.Println(coreCtx.GetHTTPStatusCode())
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return
	}
	fmt.Println(tools.Prettify(ret))
}
