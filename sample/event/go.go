package main

import (
	"context"
	"fmt"
	"lost_found/core"
	"lost_found/core/constants"
	coremodel "lost_found/core/model"
	"lost_found/core/tools"
	"lost_found/event"
	"lost_found/sample/configs"
	application "lost_found/service/application/v1"
)

func main() {
	// for redis store and logrus
	// var conf = configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
	// var conf = configs.TestConfig("https://open.feishu.cn")
	var conf = configs.TestConfig(constants.DomainFeiShu)

	application.SetAppOpenEventHandler(conf, func(coreCtx *core.Context, appOpenEvent *application.AppOpenEvent) error {
		fmt.Println(coreCtx.GetRequestID())
		fmt.Println(appOpenEvent)
		fmt.Println(tools.Prettify(appOpenEvent))
		return nil
	})

	application.SetAppStatusChangeEventHandler(conf, func(coreCtx *core.Context, appStatusChangeEvent *application.AppStatusChangeEvent) error {
		fmt.Println(coreCtx.GetRequestID())
		fmt.Println(appStatusChangeEvent.Event.AppId)
		fmt.Println(appStatusChangeEvent.Event.Status)
		fmt.Println(tools.Prettify(appStatusChangeEvent))
		return nil
	})

	header := make(map[string][]string)
	// from http request header
	header["X-Request-Id"] = []string{"63278309j-yuewuyeu-7828389"}
	req := &coremodel.OapiRequest{
		Ctx:    context.Background(),
		Header: coremodel.NewOapiHeader(header),
		Body:   "{json}", // from http request body
	}
	resp := event.Handle(conf, req)
	fmt.Println(tools.Prettify(resp))
}
