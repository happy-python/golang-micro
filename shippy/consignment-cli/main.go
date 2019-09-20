package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	pb "golang-micro/shippy/consignment-service/proto/consignment"
	"io/ioutil"
	"log"
	"os"
)

func parseFile(fileName string) (consignment *pb.Consignment, err error) {
	data, err := ioutil.ReadFile(fileName)
	err = json.Unmarshal(data, &consignment)
	return
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecing file and token."))
	}

	// 获取命令行参数
	fileName := os.Args[1]
	token := os.Args[2]

	consignment, err := parseFile(fileName)
	if err != nil {
		log.Fatalf("Could not parse file: %v\n", err)
	}

	// 生成新的context，携带数据
	ctx := metadata.NewContext(context.Background(), map[string]string{"token": token})

	consignmentClient := pb.NewShippingServiceClient("go.micro.srv.consignment", client.DefaultClient)
	resp, err := consignmentClient.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}

	log.Printf("Created: %t", resp.Created)

	resp, err = consignmentClient.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v\n", err)
	}

	for _, c := range resp.Consignments {
		log.Println(c)
	}
}
