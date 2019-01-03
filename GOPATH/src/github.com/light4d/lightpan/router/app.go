package router

import (
	"github.com/light4d/object4d/common/server"
	"net/http"
)

var staticHandler http.Handler

func app(resp http.ResponseWriter, req *http.Request) {
	staticHandler.ServeHTTP(resp, req)
}

func init() {

	staticHandler = http.FileServer(http.Dir(server.APPConfig))
}
