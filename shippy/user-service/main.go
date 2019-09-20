package main

import (
	"github.com/micro/go-micro"
	pb "golang-micro/shippy/user-service/proto/user"
	"log"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v\n", err)
	}

	db.AutoMigrate(&pb.User{})

	service := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	service.Init()

	pb.RegisterUserServiceHandler(service.Server(), &handler{repo: &Repository{db: db}})
	if err := service.Run(); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
