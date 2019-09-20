package main

import (
	"context"
	"github.com/micro/go-micro/client"
	pb "golang-micro/shippy/user-service/proto/user"
	"log"
)

func main() {
	userClient := pb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
	user := &pb.User{
		Name:     "jack",
		Email:    "jack@123.com",
		Password: "123456",
		Company:  "BBC",
	}
	// 新建用户
	resp, err := userClient.Create(context.Background(), user)
	if err != nil {
		log.Fatalf("user create err %v\n", err)
	}

	log.Printf("user create %v\n", resp.User)

	// 获取全部用户
	resp, err = userClient.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v\n", err)
	}
	for index, u := range resp.Users {
		log.Println(index, u)
	}

	// 生成token
	t, err := userClient.Auth(context.Background(), user)
	if err != nil {
		log.Fatalf("auth user err: %v\n", err)
	}

	log.Printf("Your access token is %v\n", t.Token)
}
