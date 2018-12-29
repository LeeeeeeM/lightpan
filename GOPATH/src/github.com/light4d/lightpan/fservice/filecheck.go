package fservice

import (
	"github.com/light4d/lightpan/model"
	om "github.com/light4d/object4d/model"

	"github.com/light4d/lightpan/mservice"
	"strconv"
)
//CheckUrlAccess
// 1.who want
// 2.who's

func CheckVistorAccess(uid string, f  model.File) (access bool, err error) {
	if uid == f.User {
		access = true
		return
	}
	g, b := mservice.CheckGroupExist(f.User)

	if !b {
		return
	}
	gus, err := mservice.SearchGroupuser(map[string]interface{}{
		"id":   g.ID,
		"user": uid,
	})
	if len(gus) == 1 {
		access = true
		return
	}
	err = om.NewErr("len(groupuser)=" + strconv.Itoa(len(gus)))
	return
}
