package hander

import (
	"context"
	"errors"

	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/metadata"

	"github.com/gomsa/tools/uitl"
	"github.com/gomsa/user-srv/client"
	casbinPB "github.com/gomsa/user-srv/proto/casbin"
	userPB "github.com/gomsa/user-srv/proto/user"

	pb "github.com/gomsa/user-api/proto/user"
	"github.com/gomsa/user-api/providers/redis"
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

// MobileBuild 绑定手机
func (srv *User) MobileBuild(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 通过 uuid 获取存储的验证码进行验证 req.Uuid req.Verify
	redis := redis.Client
	verify,err := redis.Get("req.Uuid").Result()
	if err != nil {
		log.Log(err)
		err = errors.New("绑定手机时,验证码未找到")
		return err
	}
	if verify != req.Verify {
		err = errors.New("验证码错误")
		return err
	}
	// meta["user_id"] 通过 meta 获取用户 id --- So this function needs token to use
	meta, _ := metadata.FromContext(ctx)
	if userID, ok := meta["user_id"]; ok {
		// 验证通过 更新用户绑定手机和用户信息 req.User
		user := &userPB.User{}
		req.User.Id = userID
		err = uitl.Data2Data(req.User, user)
		if err != nil {
			return err
		}
		userRes, err := client.User.Update(ctx, user)
		if err != nil {
			res.Valid = false
			err = errors.New("绑定手机时,更新用户信息失败")
			return err
		}
		res.Valid = userRes.Valid
	} else {
		res.Valid = false
		err = errors.New("绑定手机时,未找到用户ID")
		return err
	}
	// 返回是否绑定成功 res.Valid
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
	meta, _ := metadata.FromContext(ctx)
	user.Origin = meta["service"]
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
