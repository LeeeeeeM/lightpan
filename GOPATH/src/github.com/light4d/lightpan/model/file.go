package model

import (
	"github.com/gobestsdk/gobase/utils"
	"strings"
	"time"
)

type File struct {
	User, Folder, Name string
}

type Object4dFile struct {
	User, Folder, Name string
	Pub, Del           bool
	Version            int
	Object4d           string
	Createtime         interface{}
}

func (u *Object4dFile) FixShow() *Object4dFile {
	if u.Createtime != nil {
		u.Createtime = (u.Createtime.(time.Time)).Format(utils.DateTimeFormart)
	}

	return u
}

func ParseFile(remoteuri string) (f *File) {
	f = new(File)
	remoteuri = remoteuri[1:]
	f.User = remoteuri[:strings.Index(remoteuri, "/")]
	remoteuri = remoteuri[strings.Index(remoteuri, "/"):]
	f.Folder = remoteuri[:strings.LastIndex(remoteuri, "/")]
	f.Name = remoteuri[strings.LastIndex(remoteuri, "/")+1:]
	return
}
