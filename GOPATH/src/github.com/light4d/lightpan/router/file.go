package router

import (
	"github.com/gobestsdk/gobase/httpserver"
	"github.com/light4d/lightpan/model"

	oc "github.com/light4d/object4d/common"
	om "github.com/light4d/object4d/model"
	os "github.com/light4d/object4d/service"
	"net/http"

	ls "github.com/light4d/lightpan/common/server"

	"github.com/gobestsdk/gobase/log"
	"io/ioutil"

	"fmt"
	"github.com/light4d/lightpan/fservice"
	"github.com/light4d/object4d/common/server"
	"strings"
	"time"
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
		httpserver.Options(req, resp, "application/octet-stream", server.Appname, AccessControlAllowMethods())
	}
}

func file_get(resp http.ResponseWriter, req *http.Request) {
	result := om.CommonResp{}
	uid := getuid(req)
	f := model.ParseFile(req.URL.Path)
	access, err := fservice.CheckVistorAccess(uid, f)

	if err != nil {
		result.Code = -1
		result.Error = err.Error()

		log.Warn(log.Fields{
			"err": err.Error(),

			"result": result,
		})
		Endresp(result, resp)
		return
	}
	if !access {
		result.Code = -1
		result.Error = om.NewErr("access deny")
		Endresp(result, resp)
		return
	}
	f4d, folder, err := fservice.GetFile(f)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()

		log.Warn(log.Fields{
			"err":    err.Error(),
			"result": result,
		})
		Endresp(result, resp)
		return
	}
	if f4d != nil {
		BridgePost(f4d, req, resp)

		return
	}
	if folder != nil {
		result.Result = folder
		Endresp(result, resp)
		return
	}
}
func file_post(resp http.ResponseWriter, req *http.Request) {
	result := om.CommonResp{}
	uid := getuid(req)
	if uid == "" {
		result.Code = -1
		result.Error = om.NewErr("who are you")
		log.Info(log.Fields{
			"result": result,
		})
		Endresp(result, resp)
		return
	}
	query := httpserver.Getfilter(req)
	f := model.ParseFile(req.URL.Path)

	access, err := fservice.CheckVistorAccess(uid, f)

	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		log.Warn(log.Fields{
			"access": access,
			"err":    err.Error(),
		})
		Endresp(result, resp)
		return
	}
	if !access {
		result.Code = -1
		result.Error = om.NewErr("access deny")
		Endresp(result, resp)
		return
	}
	newobj4d, err := os.GetLocation(req)
	if err != nil {
		result = om.CommonResp{
			Error: err.Error(),
			Code:  -1,
		}
		log.Warn(log.Fields{
			"error": err.Error(),
		})
		Endresp(result, resp)
		return
	}
	f4d, folder, err := fservice.GetFile(f)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		log.Warn(log.Fields{
			"err": err.Error(),
		})
		Endresp(result, resp)
		return
	}
	if folder != nil {
		result = om.CommonResp{
			Error: om.NewErr("path already exist  a folder"),
			Code:  -1,
		}
		log.Warn(log.Fields{
			"error": err.Error(),
		})
		Endresp(result, resp)
		return
	}
	if f4d == nil {
		pub := false
		pub_, _ := query["pub"]
		if pub_ != nil {
			switch pub_.(type) {
			case bool:
				pub = pub_.(bool)
			}
		}

		f4d = &model.Object4dFile{
			User:       f.User,
			Path:       f.Path,
			Pub:        pub,
			Createtime: time.Now(),
			Version:    0,
			Object4d:   newobj4d.Url(),
		}
		err = fservice.NewFileRecord(*f4d)
		if err != nil {
			result.Code = -1
			result.Error = err.Error()
			log.Warn(log.Fields{
				"err": err.Error(),
			})
			Endresp(result, resp)
			return
		}
	} else {
		err = fservice.Tobeold(*f4d)
		if err != nil {
			result.Code = -1
			result.Error = err.Error()
			log.Warn(log.Fields{
				"err": err.Error(),
			})
			Endresp(result, resp)
			return
		}
		f4d.Createtime = time.Now()
		f4d.Object4d = newobj4d.Url()
		err = fservice.NewFileRecord(*f4d)
		if err != nil {
			result.Code = -1
			result.Error = err.Error()
			log.Warn(log.Fields{
				"err": err.Error(),
			})
			Endresp(result, resp)
			return
		}
	}
	{
		if err = BridgePost(f4d, req, resp); err != nil {
			result.Code = -1
			result.Error = err.Error()
			log.Warn(log.Fields{
				"err": err.Error(),
			})
			Endresp(result, resp)
		}

	}

}

//Bridge 转发http流量
func BridgePost(f4d *model.Object4dFile, oreq *http.Request, resp http.ResponseWriter) (err error) {
	client := &http.Client{}
	url := "http://" + ls.APPConfig.RandomElement() + "/" + f4d.Object4d
	req, err := http.NewRequest(oreq.Method, url, oreq.Body)
	if oreq.Method == "POST" {
		req.Header.Add("Content-Type", "application/octet-stream")
		req.Header.Add(oc.Ctype, ContentType(f4d.Path))
	}

	redrictresp, err := client.Do(req)

	if err != nil {
		return
	}
	fmt.Println(redrictresp.Header)
	for k, v := range redrictresp.Header {

		resp.Header().Set(k, strings.Join(v, ","))
	}

	rrbs, _ := ioutil.ReadAll(redrictresp.Body)
	resp.Write(rrbs)
	return
}
func file_delete(resp http.ResponseWriter, req *http.Request) {

}
