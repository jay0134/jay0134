FROM golang:latest

# 创建一个app目录
RUN mkdir /app
# 切换工作目录
WORKDIR /app
# 将当前的目录的文件都拷贝到
COPY . /app

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

# 将我们的代码编译成二进制可执行文件  可执行文件名为 httpserver
RUN go build -o httpserver ./httpserver.go

# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令
CMD ["/app/httpserver"]



