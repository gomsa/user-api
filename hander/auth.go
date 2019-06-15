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
	authRes, err := client.Auth.Auth(ctx, user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(authRes, res)
	if err != nil {
		return err
	}
	return err
}

// Logout 登录退出
// 后期开发
// token 增加前期认证根据根据存储的 token 进行认证 和过期查询
// 通过删除 存储 token 来登出
// 通过存储多种类型的 token 来实现多端登录
// token id type type=登录类型
// 过期时间默认还是在 jwt token 中存储
func (srv *Auth) Logout(ctx context.Context, req *pb.Request, res *pb.Token) (err error) {
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
	authRes, err := client.Auth.ValidateToken(ctx, authReq)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(authRes, res)
	if err != nil {
		return err
	}
	return err
}
