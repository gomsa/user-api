package main

import (
	"context"
	"testing"

	"github.com/gomsa/user-api/hander"
	userPB "github.com/gomsa/user-api/proto/user"
)

func TestUserCreate(t *testing.T) {
	h := hander.User{}
	req := &userPB.User{
		Name:     `bvbv0111`,
		Password: `1234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas561234asdasdasdasdasdas56`,
		Mobile:   `13953186114`,
		Email:    `bvbv0a1@qq.com`,
		Origin:   `user-api`,
	}
	res := &userPB.Response{}
	err := h.Create(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
