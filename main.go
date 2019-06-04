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
	m "github.com/gomsa/user-srv/middleware"
)

func main() {
	Pata := []map[string]interface{}{
		{"service": "Users", "method": "Create", "auth": false, "policy": false, "display_name": "创建用户", "description": "创建新用户权限。"},
		{"service": "Users", "method": "Exist", "auth": false, "policy": false, "display_name": "检测用户", "description": "检测用户是否存在权限。"},
		{"service": "Users", "method": "Get", "auth": true, "policy": false, "display_name": "用户查询", "description": "查询用户信息权限。"},
		{"service": "Auth", "method": "Auth", "auth": false, "policy": false, "display_name": "用户授权", "description": "用户登录授权返回 token 权限。"},
		{"service": "Auth", "method": "ValidateToken", "auth": false, "policy": false, "display_name": "权限认证", "description": "权限相关认证权限。"},
	}
	// 设置权限
	h := m.Handler{
		P: m.Permission{
			Data: Pata,
		},
	}
	srv := k8s.NewService(
		micro.Name("user-api"),
		micro.Version("latest"),
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
