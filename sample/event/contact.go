package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	eventhttp "lost_found/event/http"
	"lost_found/sample/configs"
	contact "lost_found/service/contact/v3"
)

func main() {

	// for redis store and logrus
	// var conf = configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
	// var conf = configs.TestConfig("https://open.feishu.cn")
	var conf = configs.TestConfig(constants.DomainFeiShu)

	contact.SetDepartmentCreatedEventHandler(conf, func(ctx *core.Context, event *contact.DepartmentCreatedEvent) error {
		fmt.Println(ctx.GetRequestID())
		fmt.Println(tools.Prettify(event))
		return nil
	})

	contact.SetUserCreatedEventHandler(conf, func(ctx *core.Context, event *contact.UserCreatedEvent) error {
		fmt.Println(ctx.GetRequestID())
		fmt.Println(tools.Prettify(event))
		return nil
	})

	contact.SetUserUpdatedEventHandler(conf, func(ctx *core.Context, event *contact.UserUpdatedEvent) error {
		fmt.Println(ctx.GetRequestID())
		fmt.Println(tools.Prettify(event))
		return nil
	})

	g := gin.Default()
	g.POST("/webhook/event", func(context *gin.Context) {
		eventhttp.Handle(conf, context.Request, context.Writer)
	})
	err := g.Run(":8089")
	if err != nil {
		fmt.Println(err)
	}

}
