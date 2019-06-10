package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gomsa/user-api/hander"
	userPB "github.com/gomsa/user-api/proto/user"
)

func TestPermissions(t *testing.T) {
	// Sync 同步权限
	permissions, _ := json.Marshal(Conf.Permissions)
	fmt.Println(permissions)
}
func TestUserGet(t *testing.T) {
	h := hander.User{}
	req := &userPB.User{
		Username: `bvbv0111`,
		Password: `1234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas56`,
		Mobile:   `13953186114`,
		Email:    `bvbv0a1@qq.com`,
		Origin:   `user-api`,
	}
	res := &userPB.Response{}
	err := h.Get(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
