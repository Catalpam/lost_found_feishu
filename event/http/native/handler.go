package native

import (
	"lost_found/core/config"
	. "lost_found/event/http"
	"net/http"
)

func Register(path string, conf *config.Config) {
	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		Handle(conf, request, writer)
	})
}
