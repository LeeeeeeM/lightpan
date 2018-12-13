package server

import (
	"github.com/gobestsdk/gobase/httpserver"
)

var (
	M = httpserver.New("lightpan manager")
	F = httpserver.New("lightpan file system")
	O = httpserver.New("ojbect4d")
)

func Run() {

	M.SetPort(APPConfig.HttpPort)
	F.SetPort(APPConfig.FsPort)
	O.SetPort(APPConfig.Object4dPort)
	go M.Run()
	go O.Run()
	F.Run()
}

func Stop() {
	M.Stop()
	O.Stop()
	F.Stop()
}
