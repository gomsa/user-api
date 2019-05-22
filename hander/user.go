package hander

import (
	"context"
	"errors"

	"github.com/micro/go-micro/metadata"

	"github.com/gomsa/tools/uitl"
	pb "github.com/gomsa/user-api/proto/user"
	"github.com/gomsa/user-srv/client"
	userPB "github.com/gomsa/user-srv/proto/user"
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
	userRes, err := client.User.Exist(context.TODO(), user)
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
	if userId, ok := meta["user_id"]; ok {
		user := &userPB.User{
			Id: userId,
		}
		userRes, err := client.User.Get(context.TODO(), user)
		if err != nil {
			return err
		}
		err = uitl.Data2Data(userRes, res)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Empty userId")
	}
	return err
}

// List 用户列表
func (srv *User) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	user := &userPB.Request{}
	err = uitl.Data2Data(req, user)
	if err != nil {
		return err
	}
	userRes, err := client.User.List(context.TODO(), user)
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
	userRes, err := client.User.Get(context.TODO(), user)
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
	userRes, err := client.User.Create(context.TODO(), user)
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
	userRes, err := client.User.Update(context.TODO(), user)
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
	userRes, err := client.User.Delete(context.TODO(), user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}
