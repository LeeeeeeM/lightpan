package router

import (
	"github.com/light4d/lightpan/common/server"
)

func Init() {

	server.M.ServerMux.HandleFunc("/user", checktoken)
	server.M.ServerMux.HandleFunc("/group", checktoken)
	server.M.ServerMux.HandleFunc("/group/owner", checktoken)
	server.M.ServerMux.HandleFunc("/group/user", checktoken)
	server.M.ServerMux.HandleFunc("/login", login)
	server.M.ServerMux.HandleFunc("/regist", user_post)

	server.F.ServerMux.HandleFunc("/", file)
	server.A.ServerMux.HandleFunc("/", app)

}
