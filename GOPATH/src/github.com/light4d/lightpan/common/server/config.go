package server

import (
	"encoding/json"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/lightpan/common/config"
	"io/ioutil"
)

var (
	APPConfig config.Config = config.Config{
		HttpPort: 9003,
		FsPort:   9002,
		Mysql:    "",
		Object4d: []string{},
		Dist:     "dist",
	}
)

func ParseConfig(configfilepath string) error {
	data, err := ioutil.ReadFile(configfilepath)
	if err != nil {
		log.Fatal(log.Fields{"error": err, "app": "config file read "})
		return err
	}

	err = json.Unmarshal([]byte(data), &APPConfig)
	if err != nil {
		log.Fatal(log.Fields{"error": err, "app": "config file parse "})
		return err
	}
	log.Info(log.Fields{"app": "config file", "config": APPConfig})
	return nil
}
