package hander

import (
	"context"
	"errors"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/util/log"

	client "github.com/gomsa/tools/k8s/client"

	casbinPB "github.com/gomsa/user-api/proto/casbin"
	pb "github.com/gomsa/user-api/proto/user"
	"github.com/gomsa/user-api/providers/redis"
)

// User 用户结构
type User struct {
	ServiceName string
}

// Exist 用户是否存在
func (srv *User) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Users.Exist", req, res)
}

// MobileBuild 绑定手机
func (srv *User) MobileBuild(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 通过 uuid 获取存储的验证码进行验证 req.Uuid req.Verify
	redis := redis.NewClient()
	verify, err := redis.Get(req.Uuid).Result()
	if err != nil {
		log.Log(err)
		err = errors.New("绑定手机时,验证码超时!")
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
		req.User.Id = userID
		err = client.Call(ctx, srv.ServiceName, "Users.Update", req, res)
		if err != nil {
			err = errors.New("绑定手机时,更新用户信息失败")
			return err
		}
	} else {
		res.Valid = false
		err = errors.New("绑定手机时,未找到用户ID")
		return err
	}
	// 返回是否绑定成功 res.Valid
	return err
}

// SelfUpdate 用户通过 token 自己更新数据 只可以更改 用户名、昵称、头像
func (srv *User) SelfUpdate(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// meta["user_id"] 通过 meta 获取用户 id --- So this function needs token to use
	meta, _ := metadata.FromContext(ctx)
	if userID, ok := meta["user_id"]; ok {
		req.User.Id = userID
		err = client.Call(ctx, srv.ServiceName, "Users.Update", req, res)
		if err != nil {
			err = errors.New("更新用户信息失败")
			return err
		}
	} else {
		err = errors.New("更新用户失败,未找到用户ID")
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
		// 获取用户信息
		if req.User == nil {
			req.User = &pb.User{}
		}
		req.User.Id = userID
		err = client.Call(ctx, srv.ServiceName, "Users.Get", req, res)
		if err != nil {
			return err
		}
		if res.User != nil {
			res.User.Password = ""
		}
		// 获取角色信息
		rolesRes := &casbinPB.Response{}
		err = client.Call(ctx, srv.ServiceName, "Casbin.GetRoles", &casbinPB.Request{
			Label: userID,
		}, rolesRes)
		if err != nil {
			return err
		}
		// 获取前端权限
		if rolesRes.Roles != nil {
			frontPermit := []string{}
			for _, roles := range rolesRes.Roles {
				frontPermitRes := &casbinPB.Response{}
				err = client.Call(ctx, srv.ServiceName, "Casbin.GetRoles", &casbinPB.Request{
					Label: roles,
				}, frontPermitRes)
				if err != nil {
					return err
				}
				if frontPermitRes.Roles != nil {
					frontPermit = append(frontPermit, frontPermitRes.Roles...)
				}
			}
			res.FrontPermits = frontPermit
			res.Roles = rolesRes.Roles
		}
	} else {
		return errors.New("Empty userID")
	}
	return err
}

// List 用户列表
func (srv *User) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Users.List", req, res)
}

// Get 获取用户
func (srv *User) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	err = client.Call(ctx, srv.ServiceName, "Users.Get", req, res)
	if res.User != nil {
		res.User.Password = ""
	}
	return err
}

// Create 创建用户
func (srv *User) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	meta, _ := metadata.FromContext(ctx)
	req.User.Origin = meta["service"]
	err = client.Call(ctx, srv.ServiceName, "Users.Create", req, res)
	if res.User != nil {
		res.User.Password = ""
	}
	return err
}

// Update 更新用户
func (srv *User) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Users.Update", req, res)
}

// Delete 删除用户
func (srv *User) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Users.Delete", &pb.User{
		Id: req.User.Id,
	}, res)
}
