package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/event"
	eventhttp "lost_found/event/http"
	"lost_found/sample/configs"
	application "lost_found/service/application/v1"
)

// for redis store and logrus
// var conf = configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// var conf = configs.TestConfig("https://open.feishu.cn")
var conf = configs.TestConfig(constants.DomainFeiShu)

func main() {

	application.SetAppOpenEventHandler(conf, func(ctx *core.Context, appOpenEvent *application.AppOpenEvent) error {
		fmt.Println(ctx.GetRequestID())
		fmt.Println(appOpenEvent)
		fmt.Println(tools.Prettify(appOpenEvent))
		return nil
	})

	/*
		application.SetAppStatusChangeEventHandler(conf, func(ctx *core.Context, appStatusChangeEvent *application.AppStatusChangeEvent) error {
			fmt.Println(ctx.GetRequestID())
			fmt.Println(appStatusChangeEvent.Event.AppId)
			fmt.Println(appStatusChangeEvent.Event.Status)
			fmt.Println(tools.Prettify(appStatusChangeEvent))
			return nil
		})
	*/
	event.SetTypeCallback(conf, "app_status_change", func(ctx *core.Context, event map[string]interface{}) error {
		fmt.Println(ctx.GetRequestID())
		fmt.Println(tools.Prettify(event))
		data := event["event"].(map[string]interface{})
		fmt.Println(tools.Prettify(data))
		return nil
	})

	application.SetAppUninstalledEventHandler(conf, func(ctx *core.Context, appUninstalledEvent *application.AppUninstalledEvent) error {
		fmt.Println(ctx.GetRequestID())
		fmt.Println(tools.Prettify(appUninstalledEvent))
		return nil
	})

	g := gin.Default()

	g.POST("/webhook/event", func(context *gin.Context) {
		eventhttp.Handle(conf, context.Request, context.Writer)
	})
	err := g.Run(":8089")
	if err != nil {
		panic(err)
	}
}
