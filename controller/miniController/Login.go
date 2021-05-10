package miniController

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/api/core/response"
	"lost_found/common"
	"lost_found/core"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/dbModel"
	authen "lost_found/service/authen/mini"
	"net/http"
)

var conf = configs.FeishuConfig(constants.DomainFeiShu)
var authenService = authen.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func SetCookies(ctx *gin.Context) {
	db := common.GetDB()
	code, _ := ctx.GetPostForm("code")
	//从飞书服务器请求用户信息
	Session, OpenId := GetAccessToken(code)
	//判断Code是否合法
	if OpenId == ""{
		ctx.SetCookie("miniAuth", "InlegalCode", 3600, "/", "fengzigeng.com", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 4001,
			"data": "",
			"msg":  "Code不匹配！!",
		})
		return
	}
	var loginingUser dbModel.User
	db.Where("open_id=?",OpenId).First(&loginingUser)
	if loginingUser.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 4002,
			"data": "",
			"msg":  "没有找到您的用户信息！!",
		})
		return
	}
	//设置Cookie，返回信息
	ctx.SetCookie("miniAuth", "InlegalOpenId", 3600, "/", "fengzigeng.com", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"data": gin.H{
			"open_id": OpenId,
			"name": loginingUser.Name,
		},
		"msg":  "登陆成功!",
	})
	println("---------------------SessionKey:"+Session)
	println("---------------------Open_Id:"+OpenId)
}

func GetAccessToken(code string) (sessionKey string, open_id string){
	ctx := context.Background()
	coreCtx := core.WrapContext(ctx)
	body := &authen.AuthenAccessTokenReqBody{
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
		return "",""
	}
	fmt.Println("-------result/------- \r\n "+string(tools.Prettify(result)))
	return result.SessionKey,result.OpenId
}

