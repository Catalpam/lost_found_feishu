package http

import (
	"lost_found/core/config"
	coremodel "lost_found/core/model"
	"lost_found/event"
	"net/http"
)

func Handle(conf *config.Config, request *http.Request, response http.ResponseWriter) {
	req, err := coremodel.ToOapiRequest(request)
	if err != nil {
		err = coremodel.NewOapiResponseOfErr(err).WriteTo(response)
		if err != nil {
			conf.GetLogger().Error(req.Ctx, err)
		}
		return
	}
	err = event.Handle(conf, req).WriteTo(response)
	if err != nil {
		conf.GetLogger().Error(req.Ctx, err)
	}
}
