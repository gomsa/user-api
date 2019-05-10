package client

import (
	"os"

	_ "github.com/gomsa/tools/client/k8s"
	"github.com/micro/go-micro/client"

	authPB "github.com/gomsa/user-srv/proto/auth"
	userPB "github.com/gomsa/user-srv/proto/user"
)

var (
	// User 用户客户端
	User userPB.UsersClient
	// Auth 认证客户端
	Auth authPB.AuthClient
)

func init() {
	userSrvName := os.Getenv("USER_SRV_NAME")

	User = userPB.NewUsersClient(userSrvName, client.DefaultClient)
	Auth = authPB.NewAuthClient(userSrvName, client.DefaultClient)
}
