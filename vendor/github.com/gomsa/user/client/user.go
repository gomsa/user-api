package client

import (
	"context"
	"fmt"

	"github.com/gomsa/tools/config"
	client "github.com/gomsa/tools/k8s/client"
	"github.com/micro/go-micro/util/log"

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
			permission := permissionPB.Permission{}
			permission.Service = p.Service
			permission.Method = p.Method
			permission.Name = p.Name
			permission.Description = p.Description

			req := &permissionPB.Request{
				Permission: &permission,
			}
			log.Log(req, 1)
			fmt.Println(req, 2)
			err := client.Call(context.TODO(), srv.ServiceName, "Permissions.UpdateOrCreate", req, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
