package router

import (
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/lightpan/common/server"
	"net/http"
)

var staticHandler http.Handler

func app(resp http.ResponseWriter, req *http.Request) {
	staticHandler.ServeHTTP(resp, req)
}

func appinit() {
	log.Info(log.Fields{
		"app-static": server.APPConfig.Dist,
	})
	staticHandler = http.FileServer(http.Dir(server.APPConfig.Dist))
}
