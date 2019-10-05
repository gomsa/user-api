package client

import (
	"context"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/server"

	cli "github.com/micro/go-micro/client/grpc"
	srv "github.com/micro/go-micro/server/grpc"
	bkr "github.com/micro/go-plugins/broker/grpc"
)

var (
	// DefaultClient 默认客户端
	DefaultClient client.Client
)

func init() {
	broker.DefaultBroker = bkr.NewBroker()
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()
	cmd.Init()

	DefaultClient = client.DefaultClient
}

// Call 使用默认客户端对服务进行同步调用
func Call(ctx context.Context, service string, endpoint string, req interface{}, rsp interface{}, opts ...client.CallOption) error {
	request := DefaultClient.NewRequest(service, endpoint, req)
	err := DefaultClient.Call(ctx, request, rsp, opts...)
	return err
}
