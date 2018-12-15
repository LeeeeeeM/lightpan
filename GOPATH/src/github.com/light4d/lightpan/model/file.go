package model

import (
	om "github.com/light4d/object4d/model"
)

type File struct {
	User, Folder, Name string
	Pub                bool
}

type Object4dFile struct {
	File
	om.Object4d
}
