package main

import (
	"context"
	"github.com/micro/go-micro/client"
	pb "golang-micro/shippy/vessel-service/proto/vessel"
	"log"
)

func main() {
	vesselClient := pb.NewVesselServiceClient("go.micro.srv.vessel", client.DefaultClient)
	vessel := &pb.Vessel{
		Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500,
	}
	resp, err := vesselClient.Create(context.Background(), vessel)
	if err != nil {
		log.Fatalf("vessel create err %v\n", err)
	}

	log.Printf("vessel create %v\n", resp.Created)
}
