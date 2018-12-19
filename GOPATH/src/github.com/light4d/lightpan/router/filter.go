package router

import (
	"github.com/light4d/lightpan/service"
	"github.com/light4d/object4d/model"

	"encoding/json"
	"errors"
	"github.com/gobestsdk/gobase/httpserver"
	"github.com/gobestsdk/gobase/log"
	"net/http"
	"strings"
)

func getuid(req *http.Request) string {
	c, err := req.Cookie("token")
	if err == nil && c != nil {
		return service.Checktoken(c.Value)
	}

	t := req.Header.Get("token")
	if t != "" {
		return service.Checktoken(t)
	}
	return ""
}
func AccessControlAllowMethods() string {
	var method = []string{
		http.MethodGet,
		http.MethodPost,
	}
	return strings.Join(method, ",")
}

func checktoken(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	result := model.CommonResp{}
	c, err := req.Cookie("token")
	if err != nil {
		result.Code = -2
		result.Error = "nedd token"
		Endresp(result, resp)
		return
	}
	if service.Checktoken(c.Value) == "" {
		result.Code = -2
		result.Error = errors.New("need token")
		Endresp(result, resp)
		return
	} else {
		switch req.URL.Path {
		case "/user":
			user(resp, req)
		case "/group":
			group(resp, req)
		case "/group/owner":
			group_setowner(resp, req)
		case "/group/user":
			group_user(resp, req)
		}
		return
	}
}
func Endresp(result model.CommonResp, resp http.ResponseWriter) {
	log.Info(log.Fields{
		"resp": result,
	})
	httpserver.Header(resp, "application/json", AccessControlAllowMethods())

	r, _ := json.Marshal(result)
	resp.Write(r)
}
