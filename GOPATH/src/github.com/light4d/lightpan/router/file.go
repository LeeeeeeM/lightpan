package router

import (
	"github.com/gobestsdk/gobase/httpserver"
	"github.com/light4d/lightpan/model"
	"net/http"
)

func file(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		file_get(resp, req)
	case http.MethodPost:
		file_post(resp, req)
	case http.MethodDelete:
		file_delete(resp, req)
	default:
		httpserver.Options(req, resp, "application/octet-stream", AccessControlAllowMethods())
	}
}

func file_get(resp http.ResponseWriter, req *http.Request) {
	uid := getuid(req)
	f := model.ParseFile(req.RequestURI)
	//service.GetObject(uid)
}
func file_post(resp http.ResponseWriter, req *http.Request) {

}
func file_delete(resp http.ResponseWriter, req *http.Request) {

}
