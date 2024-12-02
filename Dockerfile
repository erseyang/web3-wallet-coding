
FROM golang:1.22.8 AS builder

WORKDIR /app
COPY . .

# 编译应用程序
RUN go build -o myapp

# 构建阶段2
FROM alpine:latest

# 复制编译后的应用程序
COPY --from=builder /app/myapp /usr/local/bin/

# 设置工作目录
WORKDIR /usr/local/bin

# 容器启动时运行的命令
CMD ["myapp"]