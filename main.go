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
	p_data := []map[string]interface{}{
		{"service":"user-api",	"method":"Auth.Auth",			"auth":false,	"policy":false,	"display_name": "用户授权",		"description": "用户登录授权返回 token 权限。"},
		{"service":"user-api",	"method":"Auth.ValidateToken",	"auth":false,	"policy":false,	"display_name": "权限认证",		"description": "权限相关认证权限。"},
		{"service":"user-api",	"method":"Users.Exist",			"auth":false,	"policy":false,	"display_name": "检测用户",		"description": "检测用户是否存在权限。"},
		{"service":"user-api",	"method":"Users.Info",			"auth":true,	"policy":false,	"display_name": "检测用户",		"description": "检测用户是否存在权限。"},
		{"service":"user-api",	"method":"Users.Get",			"auth":true,	"policy":true,	"display_name": "用户查询",		"description": "查询用户信息权限。"},
		{"service":"user-api",	"method":"Users.Create",		"auth":false,	"policy":false,	"display_name": "创建用户",		"description": "创建新用户权限。"},
		{"service":"user-api",	"method":"Users.Update",		"auth":true,	"policy":true,	"display_name": "更新用户",		"description": "更新用户信息。"},
		{"service":"user-api",	"method":"Users.Delete",		"auth":true,	"policy":true,	"display_name": "删除用户",		"description": "删除用户。"},
	}
	// 设置权限
	h := m.Handler{
		Permissions: p_data
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
