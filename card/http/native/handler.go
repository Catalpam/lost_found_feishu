package native

import (
	. "lost_found/card/http"
	"lost_found/core/config"
	"net/http"
)

func Register(path string, conf *config.Config) {
	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		Handle(conf, request, writer)
	})
}
