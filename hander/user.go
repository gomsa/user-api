package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/user-api/proto/user"

	"github.com/gomsa/user-api/client"
	userPB "github.com/gomsa/user-srv/proto/user"
)

// User 用户结构
type User struct {
}

// Create 创建用户
func (srv *User) Create(ctx context.Context, req *pb.User, res *pb.Response) (err error) {

	reqq := &userPB.User{
		Name: `bvbv011`,
	}
	aa, err := client.User.Get(context.TODO(), reqq)
	fmt.Println(aa, err)
	return err
}

// IsExist 用户是否存在
func (srv *User) IsExist(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	fmt.Println(req, res)
	return err
}
