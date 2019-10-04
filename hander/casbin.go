package hander

import (
	"context"

	client "github.com/gomsa/tools/k8s/client"

	pb "github.com/gomsa/user-api/proto/casbin"
)

// Casbin 授权服务处理
type Casbin struct {
	ServiceName string
}

// AddPermission 增加权限
func (srv *Casbin) AddPermission(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Casbin.AddPermission", req, res)
}

// DeletePermissions 删除角色全部权限
func (srv *Casbin) DeletePermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Casbin.DeletePermissions", req, res)
}

// UpdatePermissions 更新角色权限
func (srv *Casbin) UpdatePermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Casbin.UpdatePermissions", req, res)
}

// GetPermissions 获取指定角色全部权限
func (srv *Casbin) GetPermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Casbin.GetPermissions", req, res)
}

// AddRole 增加权限
func (srv *Casbin) AddRole(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Casbin.AddRole", req, res)
}

// DeleteRoles 删除角色全部权限
func (srv *Casbin) DeleteRoles(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Casbin.DeleteRoles", req, res)
}

// UpdateRoles 更新角色权限
func (srv *Casbin) UpdateRoles(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Casbin.UpdateRoles", req, res)
}

// GetRoles 获取指定角色全部权限
func (srv *Casbin) GetRoles(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Casbin.GetRoles", req, res)
}
