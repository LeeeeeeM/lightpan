package service

import (
	"github.com/light4d/lightpan/model"

	"errors"
	"strconv"
)

func CheckUrlAccess(uid string, f *model.File) (access bool, f4d *model.Object4dFile, err error) {
	if uid == f.User {
		access = true
		return
	}
	f4d, err = GetFile(*f)
	if err != nil {
		return
	}

	if f4d != nil && f4d.Pub {
		access = true
		return
	}
	g, b := CheckGroupExist(f.User)

	if !b {
		return
	}
	gus, err := SearchGroupuser(map[string]interface{}{
		"id":   g.ID,
		"user": uid,
	})
	if len(gus) == 1 {
		access = true
		return
	}
	err = errors.New("len(groupuser)=" + strconv.Itoa(len(gus)))
	return
}
