g:
	git add .
	git commit -m"自动提交 git 代码"
	git push

dev:
	make build && make run

build:
	protoc -I . --micro_out=. --gogofaster_out=. proto/user/user.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/auth/auth.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/frontPermit/frontPermit.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/permission/permission.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/role/role.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/casbin/casbin.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/health/health.proto
