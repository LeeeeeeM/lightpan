package server

import (
	"github.com/gobestsdk/gobase/httpserver"
)

const Appname = "lightpan"

var (
	M = httpserver.New("lightpan_manager")
	F = httpserver.New("lightpan_file_system")
	A = httpserver.New("lightpan_app")
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
