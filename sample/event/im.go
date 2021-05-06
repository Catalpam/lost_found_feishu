package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	eventhttp "lost_found/event/http"
	"lost_found/sample/configs"
	im "lost_found/service/im/v1"
)

func main() {

	// for redis store and logrus
	// var conf = configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
	// var conf = configs.TestConfig("https://open.feishu.cn")

	var conf = configs.TestConfig(constants.DomainFeiShu)

	im.SetMessageReceiveEventHandler(conf, func(ctx *core.Context, event *im.MessageReceiveEvent) error {
		fmt.Println(ctx.GetRequestID())
		println("\n好友e hadhjadhj\n\n")
		fmt.Println(tools.Prettify(event))
		return nil
	})

	g := gin.Default()
	g.POST("/eventinit", func(context *gin.Context) {
		eventhttp.Handle(conf, context.Request, context.Writer)
	})
	err := g.Run(":1111")
	if err != nil {
		fmt.Println(err)
	}

}
