package model

import (
	om "github.com/light4d/object4d/model"
	"strings"
)

type File struct {
	User, Folder, Name string
}
type DBFile struct {
	File
	Pub bool
}

type Object4dFile struct {
	DBFile
	om.Object4d
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
