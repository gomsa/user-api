// # by www.bejson.com
// app: user-api
// version: latest
// permissions:
//   - {Service: "user-api", Method: Auth.Auth, Auth: false, Policy: false, DisplayName: 用户授权, Description: 用户登录授权返回 token 权限。}
//   - {Service: "user-api", Method: Auth.ValidateToken, Auth: false, Policy: false, DisplayName: 权限认证, Description: 权限相关认证权限。}
//   - {Service: "user-api", Method: Users.Exist, Auth: false, Policy: false, DisplayName: 检测用户, Description: 检测用户是否存在权限。}
//   - {Service: "user-api", Method: Users.Info, Auth: true, Policy: false, DisplayName: 检测用户, Description: 检测用户是否存在权限。}
//   - {Service: "user-api", Method: Users.Get, Auth: true, Policy: true, DisplayName: 用户查询, Description: 查询用户信息权限。}
//   - {Service: "user-api", Method: Users.Create, Auth: false, Policy: false, DisplayName: 创建用户, Description: 创建新用户权限。}
//   - {Service: "user-api", Method: Users.Update, Auth: true, Policy: true, DisplayName: 更新用户, Description: 更新用户信息。}
//   - {Service: "user-api", Method: Users.Delete, Auth: true, Policy: true, DisplayName: 删除用户, Description: 删除用户。}
package main

import (
	"github.com/gomsa/tools/config"
)

// Conf 配置
var Conf config.Config

func init() {
	Conf = config.Config{
		App:     "user-api",
		Version: "latest",
		Permissions: []config.Permission{
			{Service: "user-api", Method: "Auth.Auth", Auth: false, Policy: false, DisplayName: "用户授权", Description: "用户登录授权返回 token 权限。"},
			{Service: "user-api", Method: "Auth.ValidateToken", Auth: false, Policy: false, DisplayName: "权限认证", Description: "权限相关认证权限。"},
			{Service: "user-api", Method: "Users.Exist", Auth: false, Policy: false, DisplayName: "检测用户", Description: "检测用户是否存在权限。"},
			{Service: "user-api", Method: "Users.Info", Auth: true, Policy: false, DisplayName: "检测用户", Description: "检测用户是否存在权限。"},
			{Service: "user-api", Method: "Users.Get", Auth: true, Policy: true, DisplayName: "用户查询", Description: "查询用户信息权限。"},
			{Service: "user-api", Method: "Users.Create", Auth: false, Policy: false, DisplayName: "创建用户", Description: "创建新用户权限。"},
			{Service: "user-api", Method: "Users.Update", Auth: true, Policy: true, DisplayName: "更新用户", Description: "更新用户信息。"},
			{Service: "user-api", Method: "Users.Delete", Auth: true, Policy: true, DisplayName: "删除用户", Description: "删除用户。"},
		},
	}
}
