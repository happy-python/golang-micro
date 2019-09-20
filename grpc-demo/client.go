package main

import (
	"context"
	pb "golang-micro/grpc-demo/proto"
	"google.golang.org/grpc"
	"log"
)

const Target = ":8000"

func main() {
	conn, err := grpc.Dial(Target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v\n", err)
	}

	// 实例化客户端
	client := pb.NewUserInfoServiceClient(conn)
	req := &pb.UserRequest{
		Name: "jack",
	}

	// 调用服务
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}

	log.Printf("Recevied: %v\n", resp)
}
