package main

import (
	"github.com/micro/go-micro"
	pb "golang-micro/shippy/vessel-service/proto/vessel"
	"log"
	"os"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatalln("DB_HOST is invalid")
	}

	// 创建 MongoDB 的主会话，需在退出 main() 时候手动释放连接
	session, err := CreateSession(dbHost)
	defer session.Close()
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}

	service := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	service.Init()

	pb.RegisterVesselServiceHandler(service.Server(), &handler{session: session})

	if err := service.Run(); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
