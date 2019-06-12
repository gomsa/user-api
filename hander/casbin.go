package hander

import (
	"context"

	"github.com/gomsa/tools/uitl"
	pb "github.com/gomsa/user-api/proto/casbin"
	"github.com/gomsa/user-srv/client"
	casbinPB "github.com/gomsa/user-srv/proto/casbin"
)

// Casbin 授权服务处理
type Casbin struct {
}

// AddPermissions 增加权限
func (srv *Casbin) AddPermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	reqCasbin := &casbinPB.Request{}
	err = uitl.Data2Data(req, reqCasbin)
	if err != nil {
		return err
	}
	resCasbin, err := client.Casbin.AddPermissions(context.TODO(), reqCasbin)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(resCasbin, res)
	if err != nil {
		return err
	}
	return err
}

// DeletePermissions 删除角色全部权限
func (srv *Casbin) DeletePermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	reqCasbin := &casbinPB.Request{}
	err = uitl.Data2Data(req, reqCasbin)
	if err != nil {
		return err
	}
	resCasbin, err := client.Casbin.DeletePermissions(context.TODO(), reqCasbin)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(resCasbin, res)
	if err != nil {
		return err
	}
	return err
}

// GetPermissions 获取指定角色全部权限
func (srv *Casbin) GetPermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	reqCasbin := &casbinPB.Request{}
	err = uitl.Data2Data(req, reqCasbin)
	if err != nil {
		return err
	}
	resCasbin, err := client.Casbin.GetPermissions(context.TODO(), reqCasbin)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(resCasbin, res)
	if err != nil {
		return err
	}
	return err
}
