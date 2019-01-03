package server

import (
	"github.com/gobestsdk/gobase/httpserver"
)

const Appname = "lightpan"

var (
	M = httpserver.New("lightpan manager")
	F = httpserver.New("lightpan file system")
	A = httpserver.New("lightpan app")
)

func Run() {

	M.SetPort(APPConfig.HttpPort)
	F.SetPort(APPConfig.FsPort)
	A.SetPort(APPConfig.Appport)
	go M.Run()
	go F.Run()
	A.Run()
}

func Stop() {
	A.Stop()
	M.Stop()
	F.Stop()
}
