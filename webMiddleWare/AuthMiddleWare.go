package webMiddleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"lost_found/common"
)

func WebAuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionKey, cookieErr := ctx.Cookie("webAuth")

		if cookieErr != nil {
			//ctx.Redirect(302, "https://www.fengzigeng.com/api/weblogin")
			ctx.JSON(200, gin.H{
				"code": 4003,
				"data": "",
				"msg":  "权限不足,还未登录",
			})
			ctx.Abort()
			return
		}
		// get from Redis
		openId, err := common.RedisClient.Get("lost_found_web:" + sessionKey).Result()

		// 判断SessionKey是否存在
		if err == redis.Nil || openId == ""{
			ctx.JSON(200, gin.H{
				"code": 4003,
				"data": "",
				"msg":  "权限不足,请重新登录",
			})
			//ctx.Redirect(302, "https://www.fengzigeng.com/api/weblogin")
			ctx.Abort()
			return
		} else if err != nil {
			fmt.Println(err)
			panic(err)
		}
		ctx.Set("open_id", openId)
		fmt.Println("Web认证成功,用户的Open_Id值为：", openId)
	}
}
