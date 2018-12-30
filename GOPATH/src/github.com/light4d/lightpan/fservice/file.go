package fservice

import (
	"github.com/light4d/lightpan/model"

	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/lightpan/common/server"
	"github.com/light4d/object4d/dao"
	"strings"

	"strconv"
)

func GetFile(f model.File) (f4d *model.Object4dFile, folder *model.Folder, err error) {
	db := dao.DB(server.APPConfig.Mysql)
	f4ds := make([]model.Object4dFile, 0)
	err = db.Table("file").Where("user = ? and path = ? and version=0", f.User, f.Path).Find(&f4ds).Error

	if err != nil {
		log.Warn(log.Fields{
			"sql err": err,
		})
		return
	}
	if len(f4ds) == 1 {
		f4d = &(f4ds[0])
		return
	}
	if len(f4ds) == 0 {
		err = db.Table("file").Where("user = ? and version=0  and  path LIKE '"+f.Path+"%'", f.User).Find(&f4ds).Error
		if err != nil {
			log.Warn(log.Fields{
				"sql err": err,
			})
			return
		}
		if len(f4ds) > 0 {
			folder = new(model.Folder)
			Folderlist(f.Path, folder, f4ds...)
		}
	}
	return
}
func Folderlist(prepath string, folder *model.Folder, fs ...model.Object4dFile) {
	if folder == nil {
		panic("folder==nil")
	}
	for _, f := range fs {
		childpath := f.Path[len(prepath):]
		if strings.Contains(childpath, "/") {
			foldername := childpath[:strings.Index(childpath, "/")]
			if folder.Childfolder == nil {
				folder.Childfolder = make(map[string]*model.Folder)
			}
			folder.Childfolder[foldername] = (new(model.Folder))
			Folderlist(prepath+"/"+foldername, folder.Childfolder[foldername], f)
		} else {
			folder.Files = append(folder.Files, f)
		}
	}
	return
}
func NewFileRecord(f4d model.Object4dFile) (err error) {
	db := dao.DB(server.APPConfig.Mysql)
	err = db.Table("file").Create(f4d).Error
	if err != nil {
		log.Warn(log.Fields{
			"sql err": err,
		})
	}
	return
}

func Tobeold(f4d model.Object4dFile) (err error) {
	db := dao.DB(server.APPConfig.Mysql)
	originold := model.Object4dFile{}
	err = db.Table("file").Where("user = '" + f4d.User +
		"' and path = '" + f4d.Path + "'").Order("version desc").First(&originold).Error
	if err != nil {
		log.Warn(log.Fields{
			"sql err": err,
		})
		return
	}
	err = db.Table("file").Where("object4d = '" + f4d.Object4d +
		"' and user = '" + f4d.User +
		"' and path = '" + f4d.Path +
		"' and version = " + strconv.Itoa(f4d.Version)).Update(map[string]interface{}{
		"version": originold.Version + 1,
	}).Error
	if err != nil {
		log.Warn(log.Fields{
			"sql err": err,
		})
	}
	return
}
