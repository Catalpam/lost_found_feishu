package event

import (
	"lost_found/core"
	"lost_found/core/config"
	coremodel "lost_found/core/model"
	app "lost_found/event/app/v1"
	"lost_found/event/core/handlers"
	"lost_found/event/core/model"
	"sync"
)

var once sync.Once

func SetTypeHandler(conf *config.Config, eventType string, handler handlers.Handler) {
	handlers.SetTypeHandler(conf, eventType, handler)
}

// Deprecated, please use `SetTypeCallback`
func SetTypeHandler2(conf *config.Config, eventType string, callback func(ctx *core.Context, event map[string]interface{}) error) {
	SetTypeHandler(conf, eventType, &defaultHandler{callback: callback})
}

func SetTypeCallback(conf *config.Config, eventType string, callback func(ctx *core.Context, event map[string]interface{}) error) {
	SetTypeHandler(conf, eventType, &defaultHandler{callback: callback})
}

type defaultHandler struct {
	callback func(ctx *core.Context, event map[string]interface{}) error
}

func (h *defaultHandler) GetEvent() interface{} {
	e := make(map[string]interface{})
	return &e
}

func (h *defaultHandler) Handle(ctx *core.Context, event interface{}) error {
	e := event.(*map[string]interface{})
	return h.callback(ctx, *e)
}

func Handle(conf *config.Config, request *coremodel.OapiRequest) *coremodel.OapiResponse {
	once.Do(func() {
		app.SetAppTicketEventHandler(conf)
	})
	coreCtx := core.WrapContext(request.Ctx)
	conf.WithContext(coreCtx)
	httpEvent := &model.HTTPEvent{
		Request:  request,
		Response: &coremodel.OapiResponse{},
	}
	handlers.Handle(coreCtx, httpEvent)
	return httpEvent.Response
}
