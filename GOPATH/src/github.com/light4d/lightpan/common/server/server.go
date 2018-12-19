package server

import (
	"github.com/gobestsdk/gobase/httpserver"
)

var (
	M = httpserver.New("lightpan manager")
	F = httpserver.New("lightpan file system")
)

func Run() {

	M.SetPort(APPConfig.HttpPort)
	F.SetPort(APPConfig.FsPort)

	go M.Run()
	F.Run()
}

func Stop() {
	M.Stop()
	F.Stop()
}
