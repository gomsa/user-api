package hander

import (
	"context"

	"github.com/gomsa/tools/uitl"
	pb "github.com/gomsa/user-api/proto/frontPermit"
	"github.com/gomsa/user/client"
	casbinPB "github.com/gomsa/user/proto/casbin"
	frontPermitPB "github.com/gomsa/user/proto/frontPermit"
)

// FrontPermit 权限结构
type FrontPermit struct {
}

// UpdateOrCreate 最高权限同步前端权限列表
func (srv *FrontPermit) UpdateOrCreate(ctx context.Context, req *pb.FrontPermit, res *pb.Response) (err error) {
	// 前端权限写入数据库
	front := &frontPermitPB.FrontPermit{
		App:         req.App,
		Service:     req.Service,
		Method:      req.Method,
		Name:        req.Name,
		Description: req.Description,
	}
	frontPermitRes, err := client.FrontPermit.UpdateOrCreate(ctx, front)
	if err != nil {
		return err
	}
	// 计算包含权限 对前端权限分组增加需要的后端权限
	permissions := []*casbinPB.Permission{}
	for _, permission := range req.Permissions {
		permissions = append(permissions, &casbinPB.Permission{
			Service: permission.Service,
			Method:  permission.Method,
		})
	}
	// 更新权限
	client.Casbin.UpdatePermissions(ctx, &casbinPB.Request{
		Role:        req.App + "_" + req.Service + "_" + req.Method,
		Permissions: permissions,
	})
	res.Valid = true
	err = uitl.Data2Data(frontPermitRes, res)
	if err != nil {
		return err
	}
	return err
}

// All 权限列表
func (srv *FrontPermit) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	query := &frontPermitPB.Request{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	frontPermitRes, err := client.FrontPermit.All(ctx, query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(frontPermitRes, res)
	if err != nil {
		return err
	}
	return err
}

// List 权限列表
func (srv *FrontPermit) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	query := &frontPermitPB.ListQuery{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	frontPermitRes, err := client.FrontPermit.List(ctx, query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(frontPermitRes, res)
	if err != nil {
		return err
	}
	return err
}

// Get 获取权限
func (srv *FrontPermit) Get(ctx context.Context, req *pb.FrontPermit, res *pb.Response) (err error) {
	frontPermit := &frontPermitPB.FrontPermit{}
	err = uitl.Data2Data(req, frontPermit)
	if err != nil {
		return err
	}
	frontPermitRes, err := client.FrontPermit.Get(ctx, frontPermit)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(frontPermitRes, res)
	if err != nil {
		return err
	}
	return err
}

// Create 创建权限
func (srv *FrontPermit) Create(ctx context.Context, req *pb.FrontPermit, res *pb.Response) (err error) {
	frontPermit := &frontPermitPB.FrontPermit{}
	err = uitl.Data2Data(req, frontPermit)
	if err != nil {
		return err
	}
	frontPermitRes, err := client.FrontPermit.Create(ctx, frontPermit)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(frontPermitRes, res)
	if err != nil {
		return err
	}
	return err
}

// Update 更新权限
func (srv *FrontPermit) Update(ctx context.Context, req *pb.FrontPermit, res *pb.Response) (err error) {
	frontPermit := &frontPermitPB.FrontPermit{}
	err = uitl.Data2Data(req, frontPermit)
	if err != nil {
		return err
	}
	frontPermitRes, err := client.FrontPermit.Update(ctx, frontPermit)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(frontPermitRes, res)
	if err != nil {
		return err
	}
	return err
}

// Delete 删除权限
func (srv *FrontPermit) Delete(ctx context.Context, req *pb.FrontPermit, res *pb.Response) (err error) {
	frontPermit := &frontPermitPB.FrontPermit{
		Id: req.Id,
	}
	if err != nil {
		return err
	}
	frontPermitRes, err := client.FrontPermit.Delete(ctx, frontPermit)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(frontPermitRes, res)
	if err != nil {
		return err
	}
	return err
}
