package main

import (
	m "github.com/gomsa/user-srv/middleware"
)

// Permissions 权限设置
func Permissions() []m.Permission {
	return []m.Permission{
		{
			Service:     `user-api`,
			Method:      `Users.Create`,
			Auth:        false,
			Policy:      false,
			Name:        `创建用户`,
			Description: `创建新用户权限。`,
		},
		{
			Service:     `user-api`,
			Method:      `Users.Exist`,
			Auth:        false,
			Policy:      false,
			Name:        `检测用户`,
			Description: `检测用户是否存在权限。`,
		},
		{
			Service:     `user-api`,
			Method:      `Users.Get`,
			Auth:        true,
			Policy:      false,
			Name:        `用户查询`,
			Description: `查询用户信息权限。`,
		},
		{
			Service:     `user-api`,
			Method:      `Auth.Auth`,
			Auth:        false,
			Policy:      false,
			Name:        `用户授权`,
			Description: `用户登录授权返回 token 权限。`,
		},
		{
			Service:     `user-api`,
			Method:      `Auth.ValidateToken`,
			Auth:        false,
			Policy:      false,
			Name:        `权限认证`,
			Description: `权限相关认证权限。`,
		},
	}
}
