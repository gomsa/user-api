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
	var req PB.Request
	for _, p := range permissions {
		req.Permission = p
		res := &PB.Response{}
		err := client.Call(context.TODO(), srv.ServiceName, "Permissions.UpdateOrCreate", &req, res)
		if err != nil {
			return err
		}
	}
	return nil
}
