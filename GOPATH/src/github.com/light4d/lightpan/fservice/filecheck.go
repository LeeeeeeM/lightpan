package fservice

import (
	"github.com/light4d/lightpan/model"
	om "github.com/light4d/object4d/model"

	"github.com/light4d/lightpan/mservice"
	"strconv"
)

func CheckUrlAccess(uid string, f *model.File) (access bool, f4d *model.Object4dFile, err error) {

	f4d, err = GetFile(*f)
	if err != nil {
		return
	}
	if f4d != nil && f4d.Pub {
		access = true
		return
	}
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
