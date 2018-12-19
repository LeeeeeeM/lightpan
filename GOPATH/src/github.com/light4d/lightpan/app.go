package main

import (
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/lightpan/common/server"
	"github.com/light4d/lightpan/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.Setlogfile("lightpan.log")
	log.Info(log.Fields{"app": "exec args", "args": os.Args})

	defer func() {
		if error := recover(); error != nil {
			log.Fatal(log.Fields{"panic": error})
			exit(-1)
		}
	}()

	go func() {
		singals := make(chan os.Signal)
		signal.Notify(singals,
			syscall.SIGTERM,
			syscall.SIGINT,
			syscall.SIGKILL,
			syscall.SIGHUP,
			syscall.SIGQUIT,
		)
		<-singals
		exit(0)
	}()
	if len(os.Args) > 1 {
		server.ParseConfig(os.Args[1])
	}

	router.Init()
	server.Run()
}

func exit(status int) {
	server.Stop()
	log.Info(log.Fields{"app": status})
	os.Exit(status)
}
