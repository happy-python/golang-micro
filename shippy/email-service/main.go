package main

import (
	"context"
	"github.com/micro/go-micro"
	pb "golang-micro/shippy/user-service/proto/user"
	"log"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	// 模拟发送邮件
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}

func main() {
	service := micro.NewService(micro.Name("go.micro.srv.email"), micro.Version("latest"))

	service.Init()

	// 注册订阅
	err := micro.RegisterSubscriber(topic, service.Server(), new(Subscriber))

	if err = service.Run(); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
