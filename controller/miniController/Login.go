package miniController

import (
	"context"
	"fmt"
	"lost_found/api/core/request"
	"lost_found/api/core/response"
	"lost_found/common"
	"lost_found/core"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/dbModel"
	authen "lost_found/service/authen/mini"
	"net/http"

	"github.com/gin-gonic/gin"
)

var conf = configs.FeishuConfig(constants.DomainFeiShu)
var authenService = authen.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func SetCookies(ctx *gin.Context) {
	db := common.GetDB()
	code, _ := ctx.GetPostForm("code")
	//从飞书服务器请求用户信息
	result := GetAccessToken(code)
	SessionKey, OpenId := result.SessionKey, result.OpenId
	//判断Code是否合法
	if OpenId == "" {
		ctx.SetCookie("miniAuth", "InlegalCode", 3600, "/", "fengzigeng.com", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 4001,
			"data": "",
			"msg":  "您还未登录，请重启飞书或小程序!",
		})
		return
	}
	var loginingUser dbModel.User
	db.Where("open_id=?", OpenId).First(&loginingUser)
	if loginingUser.ID == 0 {
		println("Hot Debug: Before UserInfo(result.AccessToken)")
		ret := UserInfo(result.AccessToken)
		println("Hot Debug: After UserInfo(result.AccessToken)")
		newStudent := dbModel.User{
			Name:         ret.Name,
			StudentId:    "",
			OpenId:       ret.OpenId,
			Mobile:       ret.Mobile,
			DepartmentId: "",
			Avatar:       ret.AvatarUrl,
		}
		db.Create(&newStudent)
	}
	db.Where("open_id=?", OpenId).First(&loginingUser)

	//设置Cookie，返回信息
	ctx.SetCookie("miniAuth", SessionKey, 3600, "/", "fengzigeng.com", false, true)
	err := common.RedisClient.Set("feishu_user:"+SessionKey, OpenId, 0).Err()
	if err != nil {
		println("------------------用户信息写入Redis失败-------------------------")
		fmt.Println(err)
		panic(err)
	}
	println("Session写入成功！")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"data": gin.H{
			"open_id": OpenId,
			"name":    loginingUser.Name,
		},
		"msg": "登陆成功!",
	})
	println("---------------------SessionKey:" + SessionKey)
	println("---------------------Open_Id:" + OpenId)
}

func GetAccessToken(code string) *authen.UserAccessTokenInfo {
	ctx := context.Background()
	coreCtx := core.WrapContext(ctx)
	body := &authen.AuthenAccessTokenReqBody{
		Code: code,
	}
	reqCall := authenService.Authens.AccessToken(coreCtx, body)
	result, err := reqCall.Do()
	fmt.Println("-------coreCtx.GetRequestID/------- \r\n " + string(tools.Prettify(coreCtx.GetRequestID())))
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return nil
	}
	fmt.Println("-------result/------- \r\n " + string(tools.Prettify(result)))
	return result
}

func UserInfo(uToken string) *authen.UserInfo {
	ctx := context.Background()
	coreCtx := core.WrapContext(ctx)
	reqCall := authenService.Authens.UserInfo(coreCtx, request.SetUserAccessToken(uToken))

	result, err := reqCall.Do()
	fmt.Println(coreCtx.GetRequestID())
	fmt.Println(coreCtx.GetHTTPStatusCode())
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return nil
	}
	fmt.Println(tools.Prettify(result))
	return result
}
