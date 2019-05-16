package middleware

import (
	"context"
	"errors"
	"strings"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"

	authClient "github.com/gomsa/user-srv/client"
	authPb "github.com/gomsa/user-srv/proto/auth"
)

// Handler 处理器
// 包含一些高阶函数比如中间件常用的 token 验证等
type Handler struct {
	Permissions []Permission
}

// Wrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user-srv 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func (h *Handler) Wrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) (err error) {
		log.Log(h.Permissions, req.Service(), req.Method())
		for _, p := range h.Permissions {
			log.Log(req.Service(), req.Method(), p.Auth)
			// 访问的服务和方法匹配时验证 Auth 插件是否需要用户授权 如果需要验证则检测响应权限
			if p.Service == req.Service() && p.Service == req.Method() && p.Auth {
				log.Log(req.Service(), req.Method(), p.Auth)
				meta, ok := metadata.FromContext(ctx)
				if !ok {
					return errors.New("no auth meta-data found in request")
				}
				// Note this is now uppercase (not entirely sure why this is...)
				token := strings.Split(meta["Authorization"], "Bearer ")[1]
				// Auth here
				authResp, err := authClient.Auth.ValidateToken(context.Background(), &authPb.Request{
					Token:   token,
					Service: req.Service(),
					Method:  req.Method(),
				})
				if err != nil && authResp.Valid == false {
					return err
				}
			}
		}
		err = fn(ctx, req, resp)
		return err
	}
}