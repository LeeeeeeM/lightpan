package fservice

import (
	"github.com/light4d/lightpan/model"
	om "github.com/light4d/object4d/model"

	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/lightpan/mservice"
	"strconv"
)

//CheckUrlAccess
// 1.who want
// 2.who's

func CheckUrlAccess(uid string, f model.File) (access, pubonly bool, err error) {
	pubonly = true
	if uid == f.Who {
		access = true
		pubonly = false
		return
	}
	_, b := mservice.CheckUserExist(f.Who)
	if b {
		f4d, folder, e := GetFile(f, false)
		if e != nil {

			log.Warn(log.Fields{
				"err": err.Error(),
			})
			err = e
			return
		}
		if folder != nil {
			access = true
			return
		}
		if f4d.Pub {
			access = true
			return
		} else {
			return
		}
	}

	g, b := mservice.CheckGroupExist(f.Who)

	if !b {
		return
	}
	gus, err := mservice.SearchGroupuser(map[string]interface{}{
		"id":   g.ID,
		"user": uid,
	})
	if len(gus) == 1 {
		access = true
		pubonly = false
		return
	}
	err = om.NewErr("len(groupuser)=" + strconv.Itoa(len(gus)))
	return
}

func CheckWriteAccess(uid string, f model.File) (access, pubonly bool, err error) {
	pubonly = true
	if uid == f.Who {
		access = true
		pubonly = false
		return
	}
	_, b := mservice.CheckUserExist(f.Who)
	if b {
		f4d, folder, e := GetFile(f, false)
		if e != nil {

			log.Warn(log.Fields{
				"err": err.Error(),
			})
			err = e
			return
		}
		if folder != nil {
			access = true
			return
		}
		if f4d.Pub {
			access = true
			return
		} else {
			return
		}
	}

	g, b := mservice.CheckGroupExist(f.Who)

	if !b {
		return
	}
	gus, err := mservice.SearchGroupuser(map[string]interface{}{
		"id":   g.ID,
		"user": uid,
	})
	if len(gus) == 1 {
		access = true
		pubonly = false
		return
	}
	err = om.NewErr("len(groupuser)=" + strconv.Itoa(len(gus)))
	return
}
