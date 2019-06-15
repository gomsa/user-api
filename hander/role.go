package hander

import (
	"context"

	"github.com/gomsa/tools/uitl"
	pb "github.com/gomsa/user-api/proto/role"
	"github.com/gomsa/user-srv/client"
	rolePB "github.com/gomsa/user-srv/proto/role"
)

// Role 角色结构
type Role struct {
}

// All 权限列表
func (srv *Role) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	query := &rolePB.Request{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	roleRes, err := client.Role.All(ctx, query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(roleRes, res)
	if err != nil {
		return err
	}
	return err
}

// List 角色列表
func (srv *Role) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	query := &rolePB.ListQuery{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	roleRes, err := client.Role.List(ctx, query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(roleRes, res)
	if err != nil {
		return err
	}
	return err
}

// Get 获取角色
func (srv *Role) Get(ctx context.Context, req *pb.Role, res *pb.Response) (err error) {
	role := &rolePB.Role{}
	err = uitl.Data2Data(req, role)
	if err != nil {
		return err
	}
	roleRes, err := client.Role.Get(ctx, role)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(roleRes, res)
	if err != nil {
		return err
	}
	return err
}

// Create 创建角色
func (srv *Role) Create(ctx context.Context, req *pb.Role, res *pb.Response) (err error) {
	role := &rolePB.Role{}
	err = uitl.Data2Data(req, role)
	if err != nil {
		return err
	}
	roleRes, err := client.Role.Create(ctx, role)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(roleRes, res)
	if err != nil {
		return err
	}
	return err
}

// Update 更新角色
func (srv *Role) Update(ctx context.Context, req *pb.Role, res *pb.Response) (err error) {
	role := &rolePB.Role{}
	err = uitl.Data2Data(req, role)
	if err != nil {
		return err
	}
	roleRes, err := client.Role.Update(ctx, role)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(roleRes, res)
	if err != nil {
		return err
	}
	return err
}

// Delete 删除角色
func (srv *Role) Delete(ctx context.Context, req *pb.Role, res *pb.Response) (err error) {
	role := &rolePB.Role{
		Id: req.Id,
	}
	if err != nil {
		return err
	}
	roleRes, err := client.Role.Delete(ctx, role)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(roleRes, res)
	if err != nil {
		return err
	}
	return err
}
