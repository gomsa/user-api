package hander

import (
	"context"

	client "github.com/gomsa/tools/k8s/client"
	"github.com/micro/go-micro/util/log"

	pb "github.com/gomsa/user-api/proto/auth"
)

// Auth 授权服务处理
type Auth struct {
	ServiceName string
}

// Auth 授权认证
// 返回token
func (srv *Auth) Auth(ctx context.Context, req *pb.Request, res *pb.Request) (err error) {
	log.Log(srv.ServiceName, "Auth.Auth", req, res)
	return client.Call(ctx, srv.ServiceName, "Auth.Auth", req, res)
}

// Logout 登录退出
// 后期开发
// token 增加前期认证根据根据存储的 token 进行认证 和过期查询
// 通过删除 存储 token 来登出
// 通过存储多种类型的 token 来实现多端登录
// token id type type=登录类型
// 过期时间默认还是在 jwt token 中存储
func (srv *Auth) Logout(ctx context.Context, req *pb.Request, res *pb.Request) (err error) {
	return err
}

// ValidateToken 验证 token
// 并且验证 token 所属用户相关权限
func (srv *Auth) ValidateToken(ctx context.Context, req *pb.Request, res *pb.Request) (err error) {
	return client.Call(ctx, srv.ServiceName, "Auth.ValidateToken", req, res)
}
