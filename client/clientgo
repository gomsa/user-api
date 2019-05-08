package client

import (
	"os"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	bkr "github.com/micro/go-plugins/broker/grpc"
	cli "github.com/micro/go-plugins/client/grpc"
	srv "github.com/micro/go-plugins/server/grpc"

	_ "github.com/micro/go-plugins/registry/kubernetes"
	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	_ "github.com/micro/go-plugins/selector/static"

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
	// set values for registry/selector
	os.Setenv("MICRO_REGISTRY", "kubernetes")
	os.Setenv("MICRO_SELECTOR", "static")
	// setup broker/client/server
	broker.DefaultBroker = bkr.NewBroker()
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()
	cmd.Init()

	User = userPB.NewUsersClient(userSrvName, client.DefaultClient)
	Auth = authPB.NewAuthClient(userSrvName, client.DefaultClient)
}
