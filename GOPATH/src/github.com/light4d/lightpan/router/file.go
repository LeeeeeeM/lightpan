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
	"time"

	"github.com/light4d/lightpan/fservice"
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
	access, f4d, err := fservice.CheckUrlAccess(uid, f)

	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		log.Info(log.Fields{
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
	if f4d == nil {
		result.Code = -1
		result.Error = om.NewErr("not found")
		Endresp(result, resp)
		return
	}
	server, err := RandomElement(ls.APPConfig.Object4d)
	if err != nil {
        panic(err)
	}

	http.Redirect(resp, req, "http://"+server+"/"+f4d.Object4d, 303)
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
	f := model.ParseFile(req.RequestURI)

	access, f4d, err := fservice.CheckUrlAccess(uid, f)
	log.Info(log.Fields{
		"access": access,
		"f4d":    f4d,
	})
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
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
			Folder:     f.Folder,
			Name:       f.Name,
			Pub:        pub,
			Createtime: time.Now(),
			Version:    0,
			Object4d:   newobj4d.Url(),
		}
		err = fservice.NewFileRecord(*f4d)
		if err != nil {
			result.Code = -1
			result.Error = err.Error()
			Endresp(result, resp)
			return
		}
	} else {
		f4d.Version += 1
		f4d.Createtime = time.Now()
		f4d.Object4d = newobj4d.Url()
		err = fservice.NewFileRecord(*f4d)
		if err != nil {
			result.Code = -1
			result.Error = err.Error()
			Endresp(result, resp)
			return
		}
	}
	{
		server, err := RandomElement(ls.APPConfig.Object4d)
		if err != nil {
			panic(err)
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://"+server+"/"+f4d.Object4d, req.Body)
		req.Header.Add("Content-Type", "application/octet-stream")
		req.Header.Add(oc.Ctype, ContentType(f.Name))
		redrictresp, err := client.Do(req)
		defer redrictresp.Body.Close()
		if err != nil {
			result.Code = -1
			result.Error = err.Error()
			Endresp(result, resp)
			return
		}

		//redrictresp, err := http.Post("http://"+ls.APPConfig.Object4d[0]+"/"+f4d.Object4d, "application/octet-stream", req.Body)
		//if err != nil {
		//	result.Code = -1
		//	result.Error = err.Error()
		//	Endresp(result, resp)
		//	return
		//}

		rrbs, _ := ioutil.ReadAll(redrictresp.Body)
		resp.Write(rrbs)
	}

}
func file_delete(resp http.ResponseWriter, req *http.Request) {

}
