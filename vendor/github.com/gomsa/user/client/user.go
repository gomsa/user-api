package client

import (
	"context"
	"fmt"

	"github.com/gomsa/tools/config"
	client "github.com/gomsa/tools/k8s/client"

	permissionPB "github.com/gomsa/user/proto/permission"
)

// User 用户客户端
type User struct {
	ServiceName string
}

// SyncPermission 同步权限
func (srv *User) SyncPermission(permissions []config.Permission) error {
	for _, p := range permissions {
		if p.Policy {
			fmt.Println(1)
			permission := &permissionPB.Permission{}
			permission.Service = p.Service
			permission.Method = p.Method
			permission.Name = p.Name
			permission.Description = p.Description

			req := &permissionPB.Request{
				Permission: permission,
			}
			fmt.Println(2, req)
			err := client.Call(context.TODO(), srv.ServiceName, "Permissions.UpdateOrCreate", req, nil)
			fmt.Println(3, err)
			if err != nil {
				fmt.Println(4, err)
				return err
			}
			fmt.Println(5)
		}
	}
	return nil
}
