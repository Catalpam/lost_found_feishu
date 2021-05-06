package v1

import (
	"lost_found/core"
	"lost_found/core/config"
	"lost_found/core/constants"
	"lost_found/core/store"
	"lost_found/event/core/handlers"
	"lost_found/event/core/model"
	"time"
)

type AppTicketEventData struct {
	*model.BaseEventData
	AppTicket string `json:"app_ticket"`
}

type AppTicketEvent struct {
	*model.BaseEvent
	Event *AppTicketEventData `json:"event"`
}

type AppTicketEventHandler struct {
	event *AppTicketEvent
}

func (h *AppTicketEventHandler) GetEvent() interface{} {
	h.event = &AppTicketEvent{}
	return h.event
}

func (h *AppTicketEventHandler) Handle(ctx *core.Context, event interface{}) error {
	appTicketEvent := event.(*AppTicketEvent)
	conf := config.ByCtx(ctx)
	return conf.GetStore().Put(ctx, store.AppTicketKey(appTicketEvent.Event.AppID), appTicketEvent.Event.AppTicket, time.Hour*12)
}

func SetAppTicketEventHandler(conf *config.Config) {
	if conf.GetAppSettings().AppType == constants.AppTypeInternal {
		return
	}
	handlers.SetTypeHandler(conf, "app_ticket", &AppTicketEventHandler{})
}
