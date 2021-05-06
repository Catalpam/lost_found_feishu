package handlers

import (
	"lost_found/card/model"
	"lost_found/core"
	"lost_found/core/config"
)

type Handler func(*core.Context, *model.Card) (interface{}, error)

var AppID2Handler = make(map[string]Handler)

func getHandler(conf *config.Config) (Handler, bool) {
	h, ok := AppID2Handler[conf.GetAppSettings().AppID]
	return h, ok
}
