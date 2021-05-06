package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/card"
	cardtttp "lost_found/card/http"
	"lost_found/card/model"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/sample/configs"
)

func main() {

	// for redis store and logrus
	// var conf = configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
	// var conf = configs.TestConfig("https://open.feishu.cn")
	var conf = configs.TestConfig(constants.DomainFeiShu)

	card.SetHandler(conf, func(coreCtx *core.Context, card *model.Card) (interface{}, error) {
		fmt.Println(coreCtx.GetRequestID())
		fmt.Println(tools.Prettify(card.Action))
		return nil, nil
	})

	g := gin.Default()
	g.POST("/webhook/card", func(context *gin.Context) {
		cardtttp.Handle(conf, context.Request, context.Writer)
	})
	err := g.Run(":8089")
	if err != nil {
		panic(err)
	}
}
