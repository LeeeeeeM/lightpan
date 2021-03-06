package router

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	lm "github.com/light4d/lightpan/model"
	"github.com/light4d/object4d/model"
	"net/http"

	"github.com/gobestsdk/gobase/httpserver"

	"github.com/light4d/lightpan/common/server"
	"github.com/light4d/lightpan/mservice"
)

func group(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		group_get(resp, req)
	case http.MethodPost:
		group_post(resp, req)
	case http.MethodPut:
		group_put(resp, req)
	case http.MethodDelete:
		group_delete(resp, req)
	default:
		httpserver.Options(req, resp, "application/json", server.Appname, AccessControlAllowMethods())
	}
}
func group_get(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	filter := httpserver.Getfilter(req)

	gs, err := mservice.SearchGroup(filter)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
	} else {
		result.Result = gs
	}
	Endresp(result, resp)
}
func group_post(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	group := lm.Group{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	err = json.Unmarshal(body, &group)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	uid := getuid(req)
	groupid, err := mservice.CreateGroup(uid, group)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	result.Result = struct {
		ID string
	}{ID: groupid}
	Endresp(result, resp)
}

func group_put(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	id := req.URL.Query().Get("id")
	if id == "" {
		result.Code = -1
		result.Error = errors.New("id不能为空")
		Endresp(result, resp)
		return
	}
	updater := make(map[string]interface{})
	err := httpserver.To(req, &updater)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	uid := getuid(req)
	err = mservice.UpdateGroup(uid, id, updater)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
	}
	Endresp(result, resp)
}

func group_setowner(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	id := req.URL.Query().Get("id")
	if id == "" {
		result.Code = -1
		result.Error = errors.New("id不能为空")
		Endresp(result, resp)
		return
	}
	updater := struct {
		Owner string
	}{}
	err := httpserver.To(req, &updater)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	uid := getuid(req)
	err = mservice.SetOwner(uid, updater.Owner, id)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
	}
	Endresp(result, resp)
}

func group_delete(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	uid := getuid(req)
	groupid := req.URL.Query().Get("id")

	if groupid != "" {
		err := mservice.DeleteGroup(uid, groupid)
		if err != nil {
			result.Error = err.Error()
			result.Code = -1
		}
		Endresp(result, resp)
		return
	} else {
		result.Error = errors.New("which one do you want to delete?")
		Endresp(result, resp)
		return
	}

}
