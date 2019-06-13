package main

import (
	"github.com/gomsa/tools/config"
)

// Conf 配置
var Conf config.Config = config.Config{
	App:     "user-api",
	Version: "latest",
	Permissions: []config.Permission{
		{Service: "user-api", Method: "Auth.Auth", Auth: false, Policy: false, Name: "用户授权", Description: "用户登录授权返回 token 权限。"},
		{Service: "user-api", Method: "Auth.ValidateToken", Auth: false, Policy: false, Name: "权限认证", Description: "权限相关认证权限。"},

		{Service: "user-api", Method: "Users.Exist", Auth: false, Policy: false, Name: "检测用户", Description: "检测用户是否存在权限。"},
		{Service: "user-api", Method: "Users.Info", Auth: true, Policy: false, Name: "检测用户", Description: "检测用户是否存在权限。"},
		{Service: "user-api", Method: "Users.Create", Auth: false, Policy: false, Name: "创建用户", Description: "创建新用户权限。"},
		{Service: "user-api", Method: "Users.Get", Auth: true, Policy: true, Name: "查询用户", Description: "查询用户信息权限。"},
		{Service: "user-api", Method: "Users.Update", Auth: true, Policy: true, Name: "更新用户", Description: "更新用户信息。"},
		{Service: "user-api", Method: "Users.Delete", Auth: true, Policy: true, Name: "删除用户", Description: "删除用户。"},
		{Service: "user-api", Method: "Users.List", Auth: true, Policy: true, Name: "用户列表", Description: "查询用户列表"},

		{Service: "user-api", Method: "Roles.Create", Auth: true, Policy: true, Name: "创建角色", Description: "创建新角色权限。"},
		{Service: "user-api", Method: "Roles.Delete", Auth: true, Policy: true, Name: "删除角色", Description: "删除角色。"},
		{Service: "user-api", Method: "Roles.Update", Auth: true, Policy: true, Name: "更新角色", Description: "更新角色信息。"},
		{Service: "user-api", Method: "Roles.Get", Auth: true, Policy: true, Name: "查询角色", Description: "查询角色信息权限。"},
		{Service: "user-api", Method: "Roles.List", Auth: true, Policy: true, Name: "角色列表", Description: "查询角色列表"},
		{Service: "user-api", Method: "Roles.All", Auth: true, Policy: true, Name: "全部角色", Description: "查询全部角色。"},

		{Service: "user-api", Method: "Permissions.List", Auth: true, Policy: true, Name: "权限列表", Description: "查询权限列表"},
		{Service: "user-api", Method: "Permissions.All", Auth: true, Policy: true, Name: "全部权限", Description: "查询全部权限。"},

		{Service: "user-api", Method: "Casbin.AddPermission", Auth: true, Policy: true, Name: "添加角色权限", Description: "添加角色权限"},
		{Service: "user-api", Method: "Casbin.DeletePermissions", Auth: true, Policy: true, Name: "删除角色权限", Description: "删除角色权限"},
		{Service: "user-api", Method: "Casbin.UpdatePermissions", Auth: true, Policy: true, Name: "更新角色权限", Description: "更新角色权限"},
		{Service: "user-api", Method: "Casbin.GetPermissions", Auth: true, Policy: true, Name: "获取角色权限", Description: "获取角色权限"},
	},
}
