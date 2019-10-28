package client

import (
	"context"

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
		req := &PB.Request{
			Permission: p,
		}
		err := client.Call(context.TODO(), srv.ServiceName, "Permissions.UpdateOrCreate", req, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
