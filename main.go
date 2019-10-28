package main

import (
	// 公共引入

	"os"

	k8s "github.com/micro/examples/kubernetes/go/micro"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/gomsa/user-api/hander"
	authPB "github.com/gomsa/user-api/proto/auth"
	casbinPB "github.com/gomsa/user-api/proto/casbin"
	frontPermitPB "github.com/gomsa/user-api/proto/frontPermit"
	healthPB "github.com/gomsa/user-api/proto/health"
	permissionPB "github.com/gomsa/user-api/proto/permission"
	rolePB "github.com/gomsa/user-api/proto/role"
	userPB "github.com/gomsa/user-api/proto/user"

	"github.com/gomsa/user/client"
	m "github.com/gomsa/user/middleware"
)

func main() {
	UserService := os.Getenv("USER_NAME")

	// 设置权限
	h := m.Handler{
		Permissions: Conf.Permissions,
		UserService: UserService,
	}
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
		micro.WrapHandler(h.Wrapper), //验证权限
	)
	srv.Init()

	userPB.RegisterUsersHandler(srv.Server(), &hander.User{UserService})

	authPB.RegisterAuthHandler(srv.Server(), &hander.Auth{UserService})

	frontPermitPB.RegisterFrontPermitsHandler(srv.Server(), &hander.FrontPermit{UserService})

	permissionPB.RegisterPermissionsHandler(srv.Server(), &hander.Permission{UserService})

	rolePB.RegisterRolesHandler(srv.Server(), &hander.Role{UserService})

	// 权限管理服务实现
	casbinPB.RegisterCasbinHandler(srv.Server(), &hander.Casbin{UserService})

	healthPB.RegisterHealthHandler(srv.Server(), &hander.Health{})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	// 同步权限
	user := &client.User{UserService}
	if err := user.SyncPermission(Conf.Permissions); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
