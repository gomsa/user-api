package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/micro/go-micro/metadata"

	"github.com/gomsa/user-api/hander"
	userPB "github.com/gomsa/user-api/proto/user"
)

func TestPermissions(t *testing.T) {
	// Sync 同步权限
	// permissions, _ := json.Marshal(Conf.Permissions)
	// fmt.Println(permissions)
}

func TestMobileBuild(t *testing.T) {
	h := hander.User{}
	req := &userPB.Request{
		Verify: "632541",
		Uuid:   "asfdasfasfasafsafs",
	}
	res := &userPB.Response{}
	err := h.MobileBuild(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestUserGet(t *testing.T) {
	h := hander.User{}
	req := &userPB.Request{
		User: &userPB.User{
			Username: `bvbv0111`,
			Password: `1234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas56`,
			Mobile:   `13953186114`,
			Email:    `bvbv0a1@qq.com`,
			Origin:   `user-api`,
		},
	}
	res := &userPB.Response{}
	err := h.Get(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestUserInfo(t *testing.T) {
	h := hander.User{}
	req := &userPB.Request{}
	res := &userPB.Response{}
	meta := map[string]string{
		"user_id": "fed9c9a5-c03f-4d8d-be7c-95b9f24c9299",
	}
	ctx := metadata.NewContext(context.TODO(), meta)
	err := h.Info(ctx, req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
