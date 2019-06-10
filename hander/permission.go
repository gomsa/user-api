package hander

import (
	"context"

	"github.com/gomsa/tools/uitl"
	pb "github.com/gomsa/user-api/proto/permission"
	"github.com/gomsa/user-srv/client"
	permissionPB "github.com/gomsa/user-srv/proto/permission"
)

// Permission 权限结构
type Permission struct {
}

// All 权限列表
func (srv *Permission) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	query := &permissionPB.Request{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	permissionRes, err := client.Permission.All(context.TODO(), query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(permissionRes, res)
	if err != nil {
		return err
	}
	return err
}

// List 权限列表
func (srv *Permission) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	query := &permissionPB.ListQuery{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	permissionRes, err := client.Permission.List(context.TODO(), query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(permissionRes, res)
	if err != nil {
		return err
	}
	return err
}

// Get 获取权限
func (srv *Permission) Get(ctx context.Context, req *pb.Permission, res *pb.Response) (err error) {
	permission := &permissionPB.Permission{}
	err = uitl.Data2Data(req, permission)
	if err != nil {
		return err
	}
	permissionRes, err := client.Permission.Get(context.TODO(), permission)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(permissionRes, res)
	if err != nil {
		return err
	}
	return err
}

// Create 创建权限
func (srv *Permission) Create(ctx context.Context, req *pb.Permission, res *pb.Response) (err error) {
	permission := &permissionPB.Permission{}
	err = uitl.Data2Data(req, permission)
	if err != nil {
		return err
	}
	permissionRes, err := client.Permission.Create(context.TODO(), permission)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(permissionRes, res)
	if err != nil {
		return err
	}
	return err
}

// Update 更新权限
func (srv *Permission) Update(ctx context.Context, req *pb.Permission, res *pb.Response) (err error) {
	permission := &permissionPB.Permission{}
	err = uitl.Data2Data(req, permission)
	if err != nil {
		return err
	}
	permissionRes, err := client.Permission.Update(context.TODO(), permission)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(permissionRes, res)
	if err != nil {
		return err
	}
	return err
}

// Delete 删除权限
func (srv *Permission) Delete(ctx context.Context, req *pb.Permission, res *pb.Response) (err error) {
	permission := &permissionPB.Permission{
		Id: req.Id,
	}
	if err != nil {
		return err
	}
	permissionRes, err := client.Permission.Delete(context.TODO(), permission)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(permissionRes, res)
	if err != nil {
		return err
	}
	return err
}
