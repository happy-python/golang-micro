# 初始化
    go mod init

# 安装
    go get -u google.golang.org/grpc
    go get -u github.com/golang/protobuf/protoc-gen-go

# 编译
    make build

# 总结
    在上面微服务的实现过程中，发现微服务需要自己管理服务端监听的端口，客户端连接后进行调用。
    当有很多个微服务时，对端口的管理会比较麻烦，相比 gRPC，go-micro 实现了服务发现（Service Discovery）来方便的管理微服务。
