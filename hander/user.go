package hander

import (
	"context"
	"errors"

	"github.com/micro/go-micro/metadata"

	"github.com/gomsa/tools/uitl"
	pb "github.com/gomsa/user-api/proto/user"
	"github.com/gomsa/user-srv/client"
	userPB "github.com/gomsa/user-srv/proto/user"
	casbinPB "github.com/gomsa/user-srv/proto/casbin"
)

// User 用户结构
type User struct {
}

// Exist 用户是否存在
func (srv *User) Exist(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user := &userPB.User{}
	err = uitl.Data2Data(req, user)
	if err != nil {
		return err
	}
	userRes, err := client.User.Exist(ctx, user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}

// Info 获取用户
func (srv *User) Info(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.New("no auth meta-data found in request")
	}
	if userID, ok := meta["user_id"]; ok {
		userRes, err := client.User.Get(ctx, &userPB.User{
			Id: userID,
		})
		if err != nil {
			return err
		}
		err = uitl.Data2Data(userRes, res)
		if err != nil {
			return err
		}
		casbinRes, err := client.Casbin.GetRoles(ctx, &casbinPB.Request{
			UserID: userID,
		})
		if err != nil {
			return err
		}
		res.Roles = casbinRes.Roles
	} else {
		return errors.New("Empty userID")
	}
	return err
}

// List 用户列表
func (srv *User) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	query := &userPB.ListQuery{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	userRes, err := client.User.List(ctx, query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}

// Get 获取用户
func (srv *User) Get(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user := &userPB.User{}
	err = uitl.Data2Data(req, user)
	if err != nil {
		return err
	}
	userRes, err := client.User.Get(ctx, user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}

// Create 创建用户
func (srv *User) Create(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user := &userPB.User{}
	err = uitl.Data2Data(req, user)
	if err != nil {
		return err
	}
	userRes, err := client.User.Create(ctx, user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}

// Update 更新用户
func (srv *User) Update(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user := &userPB.User{}
	err = uitl.Data2Data(req, user)
	if err != nil {
		return err
	}
	userRes, err := client.User.Update(ctx, user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}

// Delete 删除用户
func (srv *User) Delete(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user := &userPB.User{
		Id: req.Id,
	}
	if err != nil {
		return err
	}
	userRes, err := client.User.Delete(ctx, user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}
