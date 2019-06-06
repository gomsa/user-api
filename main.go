package main

import (
	// 公共引入
	"github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/gomsa/user-api/hander"
	authPB "github.com/gomsa/user-api/proto/auth"
	permissionPB "github.com/gomsa/user-api/proto/permission"
	rolePB "github.com/gomsa/user-api/proto/role"
	userPB "github.com/gomsa/user-api/proto/user"

	"github.com/gomsa/tools/config"
	m "github.com/gomsa/user-srv/middleware"
)

func main() {
	// 加载配置
	conf := &config.Config{}
	conf.LoadFile("./config.yaml")

	log.Log(conf, conf.GetPermissions())
	// 设置权限
	h := m.Handler{
		Permissions: conf.GetPermissions(),
	}
	srv := k8s.NewService(
		micro.Name(conf.GetApp()),
		micro.Version(conf.GetVersion()),
		micro.WrapHandler(h.Wrapper), //验证权限
	)
	srv.Init()

	userPB.RegisterUsersHandler(srv.Server(), &hander.User{})

	authPB.RegisterAuthHandler(srv.Server(), &hander.Auth{})

	permissionPB.RegisterPermissionsHandler(srv.Server(), &hander.Permission{})

	rolePB.RegisterRolesHandler(srv.Server(), &hander.Role{})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
