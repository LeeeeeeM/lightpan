package model

import (
	"github.com/gobestsdk/gobase/utils"
	"strings"
	"time"
)

type File struct {
	User, Path string
}

type Object4dFile struct {
	User, Path string
	Name       string `gorm:"-"`
	Isfolder   bool   `gorm:"-",json:"-"`
	Pub, Del   bool
	Version    int
	Object4d   string
	Createtime interface{}
}

type Folder struct {
	Files       []Object4dFile
	Childfolder map[string]*Folder
}

func (u *Object4dFile) FixShow() *Object4dFile {
	if u.Createtime != nil {
		u.Createtime = (u.Createtime.(time.Time)).Format(utils.DateTimeFormart)
	}

	return u
}

func ParseFile(remoteuri string) (f File) {
	f = *new(File)
	if len(remoteuri) > 1 {
		remoteuri = remoteuri[1:]
		if strings.Contains(remoteuri, "/") {
			f.User = remoteuri[:strings.Index(remoteuri, "/")]
		}
		if strings.Contains(remoteuri, "/") {
			remoteuri = remoteuri[strings.Index(remoteuri, "/"):]
			f.Path = remoteuri
		}
		return
	}

	return
}
