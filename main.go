package main

import (
	// 公共引入
	"github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/gomsa/user-api/hander"
	authPB "github.com/gomsa/user-api/proto/auth"
	userPB "github.com/gomsa/user-api/proto/user"
	m "github.com/gomsa/user-srv/middleware"
)

func main() {
	// 设置权限
	w := m.Handler{Permissions()}
	srv := k8s.NewService(
		micro.Version("latest"),
		micro.WrapHandler(w.Wrapper), //验证权限
	)
	srv.Init()

	userPB.RegisterUsersHandler(srv.Server(), &hander.User{})

	authPB.RegisterAuthHandler(srv.Server(), &hander.Auth{})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
