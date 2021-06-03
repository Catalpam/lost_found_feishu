package webController

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
	authen "lost_found/service/authen/v1"
	"math/rand"
	"net/http"
	"time"
)

var authenService = authen.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func Login302(ctx *gin.Context)  {
	ctx.Redirect(302, "https://open.feishu.cn/open-apis/authen/v1/index?redirect_uri=https%3A%2F%2Fwww.fengzigeng.com%2Fapi%2Fwebcode&app_id=cli_a00d67c8e5f8500c&state=WebLogin")
}

func SetCookies(ctx *gin.Context) {
	db := common.GetDB()
	var user dbModel.User

	state := ctx.Query("state")
	fmt.Println("-------状态State/------- \r\n "+string(state))
	//检验状态是否正常
	if state != "WebLogin" {
		return
	}
	code := ctx.Query("code")
	fmt.Println("-------验证码Code/------- \r\n "+string(code))
	OpenId := GetAccessToken(code)
	if OpenId == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": gin.H{
				"open_id": OpenId,
				"name": user.Name,
			},
			"msg":  "对不起，您没有失物招领应用的可用性，请联系飞书获取!",
		})
	}
	//设置Cookie，返回信息
	SessionKey := fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000))
	err := common.RedisClient.Set("lost_found_web:"+SessionKey, OpenId, 7200*time.Second).Err()
	if err != nil {
		println("------------------用户信息写入Redis失败-------------------------")
		fmt.Println(err)
		panic(err)
	}
	ctx.SetCookie("webAuth", SessionKey, 72000, "/", "fengzigeng.com", false, true)
	println("Web登录: Session写入Redis成功！")
	db.Where("open_id=?",OpenId).First(&user)
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": gin.H{
	//		"open_id": OpenId,
	//		"name": user.Name,
	//	},
	//	"msg":  "登陆成功!",
	//})
	println("---------------------SessionKey:"+ SessionKey)
	println("---------------------Open_Id:"+OpenId)
	ctx.Redirect(302, "https://www.fengzigeng.com/management")
}

func GetAccessToken(code string) string {
	db := common.GetDB()
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
		fmt.Println("-------web端登陆失败，返回：/------- \r\n ")
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return ""
	}
	var user dbModel.User
	db.Where("open_id=?",result.OpenId).First(&user)
	fmt.Println("-------用户查找成功！！！")
	if user.ID == 0 {
		newUser := dbModel.User{
			Name:         result.Name,
			OpenId:       result.OpenId,
			Mobile:       result.Mobile,
			Avatar:       result.AvatarUrl,
		}
		db.Create(&newUser)
		fmt.Println("-------web端登陆用户认证成功，更新用户信息：/------- \r\n "+string(tools.Prettify(result)))
	} else {
		fmt.Println("-------web端登陆用户认证成功，用户不存在，创建用户信息：/------- \r\n "+string(tools.Prettify(result)))
		db.Model(&user).Update(dbModel.User{
			Name:         result.Name,
			OpenId:       result.OpenId,
			Mobile:       result.Mobile,
			Avatar:       result.AvatarUrl,
		})
	}
	return result.OpenId
}
