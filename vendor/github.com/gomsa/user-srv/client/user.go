package client

import (
	"os"

	"github.com/gomsa/tools/k8s/client"

	authPB "github.com/gomsa/user-srv/proto/auth"
	permissionPB "github.com/gomsa/user-srv/proto/permission"
	rolePB "github.com/gomsa/user-srv/proto/role"
	userPB "github.com/gomsa/user-srv/proto/user"
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
)

func init() {
	userSrvName := os.Getenv("USER_SRV_NAME")

	User = userPB.NewUsersClient(userSrvName, client.DefaultClient)
	Auth = authPB.NewAuthClient(userSrvName, client.DefaultClient)
	Permission = permissionPB.NewPermissionsClient(userSrvName, client.DefaultClient)
	Role = rolePB.NewRolesClient(userSrvName, client.DefaultClient)
}
