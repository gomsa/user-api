# user-api 用户api
## Install 依赖安装
```
GIT_TAG="v1.3.0" # change as needed
go get -d -u github.com/golang/protobuf/protoc-gen-go
// go-micro 插件
go get github.com/micro/protoc-gen-micro
// gogofaster 插件 取消自动生成自动 XXX_
go get -u github.com/gogo/protobuf/proto
go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
go get github.com/gogo/protobuf/gogoproto
```
## 调用模块
- [user](https://github.com/gomsa/user)
## user 用户服务
- Exist     查询用户
- List      用户列表
- Get       查询用户
- Create    创建用户
- Update    更新用户
- Delete    删除用户
## auth 认证服务
- Auth              用户授权
- ValidateToken     授权认证
- (**高危-内部调用**)
- AuthById          通过用户ID获取授权