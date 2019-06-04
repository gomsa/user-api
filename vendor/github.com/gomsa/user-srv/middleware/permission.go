package middleware

import (
	"github.com/micro/go-micro/server"
)

// Permission 权限
type Permission struct {
	Data []map[string]interface{}
}

// IsAuth 检测是否需要检测用户
func (p *Permission) IsAuth(req server.Request) bool {
	for _, d := range p.Data {
		if d["service"] == req.Service() && d["method"] == req.Method() && d["auth"] == true {
			return true
		}
	}
	return false
}
