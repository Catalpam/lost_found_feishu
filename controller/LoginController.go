package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	authen "lost_found/service/authen/v1"
	"lost_found/core/configs"
	"net/http"
)

var authenService = authen.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func GetAccessToken() {
	ctx := context.Background()
	coreCtx := core.WrapContext(ctx)
	body := &authen.AuthenAccessTokenReqBody{
		GrantType: "authorization_code",
		Code:      "476Bsaz9mCDIAOmjIOjD4a",
	}
	reqCall := authenService.Authens.AccessToken(coreCtx, body)

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

func SetLoginCookie(ctx *gin.Context) {

	ctx.SetCookie("site_cookie", "cookievalue", 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "",
		"msg": "请求成功!",
	})

}
