package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/card"
	"lost_found/card/http"
	"lost_found/card/model"
	"lost_found/cardMessage"
	"lost_found/core"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/core/tools"
)

func CardEvent() gin.HandlerFunc {

	var conf = configs.FeishuConfig(constants.DomainFeiShu)

	card.SetHandler(conf, func(coreCtx *core.Context, card *model.Card) (interface{}, error) {
		fmt.Println(tools.Prettify(card.Action))

		cardMessage.SendMessage(card.OpenID, "按钮点击成功！")
		return nil, nil
	})

	return func(context *gin.Context) {
		http.Handle(conf, context.Request, context.Writer)
	}
}
