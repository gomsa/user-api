package client

import (
	"context"

	"github.com/gomsa/tools/config"
	client "github.com/gomsa/tools/k8s/client"

	permissionPB "github.com/gomsa/user/proto/permission"
)

// User 用户客户端
type User struct {
	ServiceName string
}

// SyncPermission 同步权限
func (srv *User) SyncPermission(permission []config.Permission) error {
	for _, p := range permission {
		if p.Policy {
			per := &permissionPB.Permission{}
			per.Service = p.Service
			per.Method = p.Method
			per.Name = p.Name
			per.Description = p.Description

			req := &permissionPB.Request{
				Permission: per,
			}
			err := client.Call(context.TODO(), srv.ServiceName, "Permission.UpdateOrCreate", req, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
