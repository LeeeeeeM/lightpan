package service

import (
	"github.com/light4d/lightpan/model"

	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/lightpan/common/server"
	"github.com/light4d/object4d/dao"
)

func GetFile(f model.File) (o4f *model.Object4dFile, err error) {
	db := dao.DB(server.APPConfig.Mysql)
	os := []model.Object4dFile{}
	err = db.Table("file").Where("user = ? and folder = ? and name = ?", f.User, f.Folder, f.Name).Order("createtime desc").Find(&os).Error
	if err != nil {
		log.Warn(log.Fields{
			"sql err": err,
		})
	}
	if len(os) > 0 {
		o4f = &(os[0])
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
