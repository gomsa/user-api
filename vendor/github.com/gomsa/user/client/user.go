package client

import (
	"context"
	"fmt"
	"time"

	client "github.com/gomsa/tools/k8s/client"

	PB "github.com/gomsa/user/proto/permission"
)

// User 用户客户端
type User struct {
	ServiceName string
}

// SyncPermission 同步权限
func (srv *User) SyncPermission(permissions []*PB.Permission) error {
	for _, p := range permissions {
		fmt.Println(1)
		req := &PB.Request{
			Permission: p,
		}
		fmt.Println(2)
		err := client.Call(context.TODO(), srv.ServiceName, "Permissions.UpdateOrCreate", req, nil)
		fmt.Println(3, err)
		if err != nil {
			fmt.Println(31, err)
			return err
		}
		fmt.Println(4)
		time.Sleep(1 * time.Second)
	}
	return nil
}
