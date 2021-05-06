// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
	"github.com/larksuite/oapi-sdk-go/event"
)

type AppOpenEventHandler struct {
	Fn func(*core.Context, *AppOpenEvent) error
}

func (h *AppOpenEventHandler) GetEvent() interface{} {
	return &AppOpenEvent{}
}

func (h *AppOpenEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*AppOpenEvent))
}

func SetAppOpenEventHandler(conf *config.Config, fn func(ctx *core.Context, event *AppOpenEvent) error) {
	event.SetTypeHandler(conf, "app_open", &AppOpenEventHandler{Fn: fn})
}

type AppStatusChangeEventHandler struct {
	Fn func(*core.Context, *AppStatusChangeEvent) error
}

func (h *AppStatusChangeEventHandler) GetEvent() interface{} {
	return &AppStatusChangeEvent{}
}

func (h *AppStatusChangeEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*AppStatusChangeEvent))
}

func SetAppStatusChangeEventHandler(conf *config.Config, fn func(ctx *core.Context, event *AppStatusChangeEvent) error) {
	event.SetTypeHandler(conf, "app_status_change", &AppStatusChangeEventHandler{Fn: fn})
}

type AppUninstalledEventHandler struct {
	Fn func(*core.Context, *AppUninstalledEvent) error
}

func (h *AppUninstalledEventHandler) GetEvent() interface{} {
	return &AppUninstalledEvent{}
}

func (h *AppUninstalledEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*AppUninstalledEvent))
}

func SetAppUninstalledEventHandler(conf *config.Config, fn func(ctx *core.Context, event *AppUninstalledEvent) error) {
	event.SetTypeHandler(conf, "app_uninstalled", &AppUninstalledEventHandler{Fn: fn})
}

type OrderPaidEventHandler struct {
	Fn func(*core.Context, *OrderPaidEvent) error
}

func (h *OrderPaidEventHandler) GetEvent() interface{} {
	return &OrderPaidEvent{}
}

func (h *OrderPaidEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*OrderPaidEvent))
}

func SetOrderPaidEventHandler(conf *config.Config, fn func(ctx *core.Context, event *OrderPaidEvent) error) {
	event.SetTypeHandler(conf, "order_paid", &OrderPaidEventHandler{Fn: fn})
}
