# 指定基础镜像
FROM golang:1.17-alpine AS build

# 设置工作目录
WORKDIR /app

# 复制代码到容器
COPY . .

# 下载依赖
RUN go mod download

# 构建可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app

# 生产阶段使用轻量级镜像
FROM alpine:latest

# 复制可执行文件到容器
COPY --from=build /go/bin/app /usr/local/bin/

# alpine 不包含完整时区数据库，需要单独安装，即时不安装，程序也会指定默认时区 UTC
RUN apk update && apk add tzdata

# 指定启动命令
CMD ["app"]