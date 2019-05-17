package hander

import (
	"context"

	"github.com/gomsa/tools/uitl"
	pb "github.com/gomsa/user-api/proto/auth"
	"github.com/gomsa/user-srv/client"
	authPB "github.com/gomsa/user-srv/proto/auth"
)

// Auth 授权服务处理
type Auth struct {
}

// Auth 授权认证
// 返回token
func (srv *Auth) Auth(ctx context.Context, req *pb.User, res *pb.Token) (err error) {
	user := &authPB.User{}
	err = uitl.Data2Data(req, user)
	if err != nil {
		return err
	}
	authRes, err := client.Auth.Auth(context.TODO(), user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(authRes, res)
	if err != nil {
		return err
	}
	return err
}

// ValidateToken 验证 token
// 并且验证 token 所属用户相关权限
func (srv *Auth) ValidateToken(ctx context.Context, req *pb.Request, res *pb.Token) (err error) {
	authReq := &authPB.Request{}
	err = uitl.Data2Data(req, authReq)
	if err != nil {
		return err
	}
	authRes, err := client.Auth.ValidateToken(context.TODO(), authReq)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(authRes, res)
	if err != nil {
		return err
	}
	return err
}
