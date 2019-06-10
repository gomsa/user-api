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

	"github.com/gomsa/user-srv/client"
	m "github.com/gomsa/user-srv/middleware"
)

func main() {
	// 设置权限
	h := m.Handler{
		Permissions: Conf.Permissions,
	}
	srv := k8s.NewService(
		micro.Name(Conf.App),
		micro.Version(Conf.Version),
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
	// 同步权限
	client.Permission.SyncPermission(Conf.Permissions)
	log.Log("serviser run ...")
}
