package card

import (
	"lost_found/card/handlers"
	"lost_found/card/model"
	"lost_found/core"
	"lost_found/core/config"
	coremodel "lost_found/core/model"
)

func SetHandler(conf *config.Config, handler handlers.Handler) {
	handlers.AppID2Handler[conf.GetAppSettings().AppID] = handler
}

func Handle(conf *config.Config, request *coremodel.OapiRequest) *coremodel.OapiResponse {
	coreCtx := core.WrapContext(request.Ctx)
	conf.WithContext(coreCtx)
	httpCard := &model.HTTPCard{
		Request:  request,
		Response: &coremodel.OapiResponse{},
	}
	handlers.Handle(coreCtx, httpCard)
	return httpCard.Response
}