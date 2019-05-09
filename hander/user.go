package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/user-api/proto/user"
	"github.com/micro/go-log"

	"github.com/gomsa/user-api/client"
	userPB "github.com/gomsa/user-srv/proto/user"
)

// User 用户结构
type User struct {
}

// Create 创建用户
func (srv *User) Create(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	fmt.Println(req, res)
	return err
}

// IsExist 用户是否存在
func (srv *User) IsExist(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	fmt.Println(req, res)
	return err
}

// Get 获取用户
func (srv *User) Get(ctx context.Context, req *pb.User, res *pb.Response) (err error) {

	reqq := &userPB.User{
		Name: `bvbv011`,
	}
	aa, err := client.User.Get(context.TODO(), reqq)
	fmt.Println("aaaaqq111",aa, err)
	log.Log(aa, err)
	return err
}
