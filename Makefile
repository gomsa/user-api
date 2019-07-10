g:
	git add .
	git commit -m"自动提交 git 代码"
	git push

dev:
	make build && make run

build:
	protoc -I . --go_out=plugins=micro:. proto/user/user.proto
	protoc -I . --go_out=plugins=micro:. proto/auth/auth.proto
	protoc -I . --go_out=plugins=micro:. proto/frontPermit/frontPermit.proto
	protoc -I . --go_out=plugins=micro:. proto/permission/permission.proto
	protoc -I . --go_out=plugins=micro:. proto/role/role.proto
	protoc -I . --go_out=plugins=micro:. proto/casbin/casbin.proto
