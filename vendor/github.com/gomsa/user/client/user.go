package client

import (
	"context"
	"os"

	"github.com/gomsa/tools/config"
	"github.com/gomsa/tools/k8s/client"

	authPB "github.com/gomsa/user/proto/auth"
	casbinPB "github.com/gomsa/user/proto/casbin"
	permissionPB "github.com/gomsa/user/proto/permission"
	rolePB "github.com/gomsa/user/proto/role"
	userPB "github.com/gomsa/user/proto/user"
)

var (
	// User 用户客户端
	User userPB.UsersClient
	// Auth 认证客户端
	Auth authPB.AuthClient
	// Permission 权限客户端
	Permission permissionPB.PermissionsClient
	// Role 角色客户端
	Role rolePB.RolesClient
	// Casbin 权限管理客户端
	Casbin casbinPB.CasbinClient
)

func init() {
	userSrvName := os.Getenv("USER_NAME")

	User = userPB.NewUsersClient(userSrvName, client.DefaultClient)
	Auth = authPB.NewAuthClient(userSrvName, client.DefaultClient)
	Permission = permissionPB.NewPermissionsClient(userSrvName, client.DefaultClient)
	Role = rolePB.NewRolesClient(userSrvName, client.DefaultClient)
	Casbin = casbinPB.NewCasbinClient(userSrvName, client.DefaultClient)
}

// SyncPermission 同步权限
func SyncPermission(permission []config.Permission) error {
	for _, p := range permission {
		if p.Policy {
			req := permissionPB.Permission{}
			req.Service = p.Service
			req.Method = p.Method
			req.Name = p.Name
			req.Description = p.Description
			_, err := Permission.UpdateOrCreate(context.TODO(), &req)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
