package hander

import (
	"context"
	"fmt"

	"github.com/gomsa/user-api/client"
	pb "github.com/gomsa/user-api/proto/auth"
)

// Auth 授权服务处理
type Auth struct {
}

// Auth 授权认证
// 返回token
func (srv *Auth) Auth(ctx context.Context, req *pb.User, res *pb.Token) (err error) {
	client.
	fmt.Println(req, res)
	return err
}

// ValidateToken 验证 token
// 并且验证 token 所属用户相关权限
func (srv *Auth) ValidateToken(ctx context.Context, req *pb.Request, res *pb.Token) (err error) {
	fmt.Println(req, res)
	return err
}
