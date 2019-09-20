package main

import (
	"context"
	pb "golang-micro/grpc-demo/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const Address = ":8000"

type handler struct {
}

// 实现 user.pb.go 中 UserInfoServiceServer 接口
func (h *handler) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name
	// 模拟在数据库中查找用户信息
	if name == "jack" {
		resp = &pb.UserResponse{
			Id:   1,
			Name: name,
			Age:  18,
			Tags: []string{"gopher", "pythoner"},
		}
	}
	return
}

func main() {
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	log.Printf("listen: %s\n", Address)

	server := grpc.NewServer()
	// 注册
	pb.RegisterUserInfoServiceServer(server, &handler{})
	server.Serve(listener)
}
