package miniMiddleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"lost_found/common"
	"net/http"
)

func MiniAuthMiddleWare() gin.HandlerFunc  {
		return func(ctx *gin.Context) {
		sessionKey, cookieErr := ctx.Cookie("miniAuth")
		if cookieErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
		"code": 4003,
		"data":"",
		"msg":"权限不足,还未登录，请调用tt.login登录",
	})
		ctx.Abort()
		return
	}
		// get from Redis
		openId, err := common.RedisClient.Get("feishu_user:"+sessionKey).Result()

		// 判断SessionKey是否存在
		if err == redis.Nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{
		"code": 4003,
		"data":"",
		"msg":"权限不足,Cookie不合法，请调用tt.login登录",
	})
		ctx.Abort()
		return
	} else if err != nil {
		fmt.Println(err)
		panic(err)
	}
		ctx.Set("open_id",openId)

		fmt.Println("读取成功,用户的Open_Id值为：", openId)
		println("SessionKey读取成功！")
		println("----------open_id成功写入至上下文中--------------")
		//ctx.Next()
	}
}