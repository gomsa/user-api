package hander

import (
	"context"

	client "github.com/gomsa/tools/k8s/client"

	casbinPB "github.com/gomsa/user-api/proto/casbin"
	pb "github.com/gomsa/user-api/proto/frontPermit"
)

// FrontPermit 权限结构
type FrontPermit struct {
	ServiceName string
}

// UpdateOrCreate 最高权限同步前端权限列表
func (srv *FrontPermit) UpdateOrCreate(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 前端权限写入数据库
	err = client.Call(ctx, srv.ServiceName, "FrontPermits.UpdateOrCreate", req, res)
	if err != nil {
		return err
	}
	// 计算包含权限 对前端权限分组增加需要的后端权限
	permissions := []*casbinPB.Permission{}
	for _, permission := range req.FrontPermit.Permissions {
		permissions = append(permissions, &casbinPB.Permission{
			Service: permission.Service,
			Method:  permission.Method,
		})
	}

	err = client.Call(ctx, srv.ServiceName, "Casbin.UpdatePermissions", &casbinPB.Request{
		Role:        req.FrontPermit.App + "_" + req.FrontPermit.Service + "_" + req.FrontPermit.Method,
		Permissions: permissions,
	}, nil)
	if err != nil {
		return err
	}
	res.Valid = true
	return err
}

// All 权限列表
func (srv *FrontPermit) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "FrontPermits.All", req, res)
}

// List 权限列表
func (srv *FrontPermit) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "FrontPermits.List", req, res)
}

// Get 获取权限
func (srv *FrontPermit) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "FrontPermits.Get", req, res)
}

// Create 创建权限
func (srv *FrontPermit) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "FrontPermits.Create", req, res)
}

// Update 更新权限
func (srv *FrontPermit) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "FrontPermits.Update", req, res)
}

// Delete 删除权限
func (srv *FrontPermit) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "FrontPermits.Delete", &pb.FrontPermit{
		Id: req.FrontPermit.Id,
	}, res)
}
