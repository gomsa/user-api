package hander

import (
	"context"

	"github.com/gomsa/tools/uitl"
	pb "github.com/gomsa/user-api/proto/user"
	"github.com/gomsa/user-srv/client"
	userPB "github.com/gomsa/user-srv/proto/user"
)

// User 用户结构
type User struct {
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
func (srv *User) Info(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
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
