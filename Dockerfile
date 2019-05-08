# 暂未将 Golang 集成到 docker 中
FROM alpine:latest

# 安装 证书、时区
RUN apk --no-cache add ca-certificates tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY ./bin/service /usr/local/bin/
CMD ["service"]