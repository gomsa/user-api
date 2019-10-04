package client

import (
	"context"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"

	cli "github.com/micro/go-micro/client/grpc"
)

var (
	// DefaultClient 默认客户端
	DefaultClient client.Client
)

func init() {
	client.DefaultClient = cli.NewClient()

	cmd.Init()

	DefaultClient = client.DefaultClient
}

// Call 使用默认客户端对服务进行同步调用
func Call(ctx context.Context, service string, endpoint string, req interface{}, rsp interface{}, opts ...client.CallOption) error {
	request := DefaultClient.NewRequest(service, endpoint, req)
	err := DefaultClient.Call(ctx, request, rsp, opts...)
	return err
}
