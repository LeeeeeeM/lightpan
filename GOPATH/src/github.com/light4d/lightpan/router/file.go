package router

import (
	"github.com/gobestsdk/gobase/httpserver"
	"github.com/light4d/lightpan/model"
	"github.com/light4d/lightpan/service"
	om "github.com/light4d/object4d/model"
	"net/http"

	"errors"
	ls "github.com/light4d/lightpan/common/server"
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
	result := om.CommonResp{}
	uid := getuid(req)
	f := model.ParseFile(req.RequestURI)
	access, f4d, err := service.CheckFileAccess(uid, f)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	if !access {
		result.Code = -1
		result.Error = errors.New("access deny")
		Endresp(result, resp)
		return
	}
	http.Redirect(resp, req, "http://"+ls.APPConfig.Object4d[0]+"/"+f4d.Object4d.Url(), 302)
}
func file_post(resp http.ResponseWriter, req *http.Request) {

}
func file_delete(resp http.ResponseWriter, req *http.Request) {

}
