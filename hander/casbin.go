package hander

import (
	"context"

	"github.com/gomsa/tools/uitl"
	pb "github.com/gomsa/user-api/proto/casbin"
	"github.com/gomsa/user/client"
	casbinPB "github.com/gomsa/user/proto/casbin"
)

// Casbin 授权服务处理
type Casbin struct {
}

// AddPermission 增加权限
func (srv *Casbin) AddPermission(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	reqCasbin := &casbinPB.Request{}
	err = uitl.Data2Data(req, reqCasbin)
	if err != nil {
		return err
	}
	resCasbin, err := client.Casbin.AddPermission(ctx, reqCasbin)
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
	resCasbin, err := client.Casbin.DeletePermissions(ctx, reqCasbin)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(resCasbin, res)
	if err != nil {
		return err
	}
	return err
}

// UpdatePermissions 更新角色权限
func (srv *Casbin) UpdatePermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	reqCasbin := &casbinPB.Request{}
	err = uitl.Data2Data(req, reqCasbin)
	if err != nil {
		return err
	}
	resCasbin, err := client.Casbin.UpdatePermissions(ctx, reqCasbin)
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
	resCasbin, err := client.Casbin.GetPermissions(ctx, reqCasbin)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(resCasbin, res)
	if err != nil {
		return err
	}
	return err
}

// AddRole 增加权限
func (srv *Casbin) AddRole(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	reqCasbin := &casbinPB.Request{}
	err = uitl.Data2Data(req, reqCasbin)
	if err != nil {
		return err
	}
	resCasbin, err := client.Casbin.AddRole(ctx, reqCasbin)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(resCasbin, res)
	if err != nil {
		return err
	}
	return err
}

// DeleteRoles 删除角色全部权限
func (srv *Casbin) DeleteRoles(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	reqCasbin := &casbinPB.Request{}
	err = uitl.Data2Data(req, reqCasbin)
	if err != nil {
		return err
	}
	resCasbin, err := client.Casbin.DeleteRoles(ctx, reqCasbin)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(resCasbin, res)
	if err != nil {
		return err
	}
	return err
}

// UpdateRoles 更新角色权限
func (srv *Casbin) UpdateRoles(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	reqCasbin := &casbinPB.Request{}
	err = uitl.Data2Data(req, reqCasbin)
	if err != nil {
		return err
	}
	resCasbin, err := client.Casbin.UpdateRoles(ctx, reqCasbin)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(resCasbin, res)
	if err != nil {
		return err
	}
	return err
}

// GetRoles 获取指定角色全部权限
func (srv *Casbin) GetRoles(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	reqCasbin := &casbinPB.Request{}
	err = uitl.Data2Data(req, reqCasbin)
	if err != nil {
		return err
	}
	resCasbin, err := client.Casbin.GetRoles(ctx, reqCasbin)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(resCasbin, res)
	if err != nil {
		return err
	}
	return err
}
