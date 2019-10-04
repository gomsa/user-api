package client

import (
	"context"
	"os"

	client "github.com/gomsa/tools/k8s/client"
	"github.com/gomsa/tools/config"

	permissionPB "github.com/gomsa/user/proto/permission"
)


// User 用户客户端
type User struct {
	ServiceName string
}
// SyncPermission 同步权限
func (srv *User)SyncPermission(permission []config.Permission) error {
	for _, p := range permission {
		if p.Policy {
			req := permissionPB.Permission{}
			req.Service = p.Service
			req.Method = p.Method
			req.Name = p.Name
			req.Description = p.Description
			err := client.Call(context.TODO(), srv.ServiceName, "Permission.UpdateOrCreate", req, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}