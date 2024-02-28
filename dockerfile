# 使用官方 Golang 镜像作为构建环境
FROM golang:1.18-alpine

# 设置工作目录
WORKDIR /app

# 复制 go 模块和依赖文件
COPY go.mod go.sum ./

# 下载依赖项
RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o demo_gin .

# 暴露 8080 端口
EXPOSE 8080

# 运行应用程序
CMD ["./demo_gin"]

