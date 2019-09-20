package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	pb "golang-micro/shippy/consignment-service/proto/consignment"
	userPb "golang-micro/shippy/user-service/proto/user"
	vesselPb "golang-micro/shippy/vessel-service/proto/vessel"
	"log"
	"os"
)

// AuthWrapper 是一个高阶函数，入参是"下一步"函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-cli 上下文中取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// 注意这里是大写的Token
		token := meta["Token"]

		// Auth
		userClient := userPb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		authResp, err := userClient.ValidateToken(context.Background(), &userPb.Token{
			Token: token,
		})
		log.Println("Auth Resp:", authResp.Valid)
		if err != nil {
			log.Printf("validate token err: %v\n", err)
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}

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
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	service.Init()

	vesselClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", service.Client())
	pb.RegisterShippingServiceHandler(service.Server(), &handler{session: session, vesselClient: vesselClient})

	if err := service.Run(); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
