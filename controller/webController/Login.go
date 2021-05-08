package webController

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/core/tools"
	authen "lost_found/service/authen/v1"
	"net/http"
)

var authenService = authen.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func Login302(ctx *gin.Context)  {
	ctx.Redirect(302, "https://open.feishu.cn/open-apis/authen/v1/index?redirect_uri=https%3A%2F%2Fwww.fengzigeng.com%2Fapi%2Fwebcode&app_id=cli_a00d67c8e5f8500c&state=WebLogin")
}

func SetCookies(ctx *gin.Context) {
	state := ctx.Query("state")
	fmt.Println("-------状态State/------- \r\n "+string(state))
	//检验状态是否正常
	if state != "WebLogin" {
		return
	}
	code := ctx.Query("code")
	fmt.Println("-------验证码Code/------- \r\n "+string(code))
	GetAccessToken(code)
	ctx.SetCookie("lost_found_web", "cookievalue", 3600, "/", "fengzigeng.com", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "",
		"msg": "登陆成功!",
	})
}

func GetAccessToken(code string) {
	ctx := context.Background()
	coreCtx := core.WrapContext(ctx)
	body := &authen.AuthenAccessTokenReqBody{
		GrantType: "authorization_code",
		Code: code,
	}
	reqCall := authenService.Authens.AccessToken(coreCtx, body)
	result, err := reqCall.Do()
	fmt.Println("-------coreCtx.GetRequestID/------- \r\n "+string(tools.Prettify(coreCtx.GetRequestID())))
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return
	}
	fmt.Println("-------result/------- \r\n "+string(tools.Prettify(result)))
}
