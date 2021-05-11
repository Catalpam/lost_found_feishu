package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/core"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/core/tools"
	card2 "lost_found/handler/card"
	"lost_found/handler/card/http"
	model2 "lost_found/handler/card/model"
)

func main() {

	// for redis store and logrus
	// var conf = configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
	// var conf = configs.TestConfig("https://open.feishu.cn")
	var conf = configs.FeishuConfig(constants.DomainFeiShu)

	card2.SetHandler(conf, func(coreCtx *core.Context, card *model2.Card) (interface{}, error) {
		fmt.Println(coreCtx.GetRequestID())
		fmt.Println(tools.Prettify(card.Action))
		return nil, nil
	})

	g := gin.Default()
	g.POST("/webhook/card", func(context *gin.Context) {
		http.Handle(conf, context.Request, context.Writer)
	})
	err := g.Run(":8089")
	if err != nil {
		panic(err)
	}
}
