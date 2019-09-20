package main

import (
	"context"
	pb "golang-micro/shippy/consignment-service/proto/consignment"
	vesselPb "golang-micro/shippy/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"log"
)

type handler struct {
	session      *mgo.Session
	vesselClient vesselPb.VesselServiceClient
}

func (h *handler) getRepository() IRepository {
	// 从主会话中 Clone() 出新会话处理查询
	return &Repository{
		session: h.session.Clone(),
	}
}

func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	vesselResp, err := h.vesselClient.FindAvailable(context.Background(), &vesselPb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}

	log.Printf("Found vessel: %s\n", vesselResp.Vessel.Name)

	req.VesselId = vesselResp.Vessel.Id

	err = h.getRepository().Create(req)
	if err != nil {
		log.Printf("create consignment err: %v\n", err)
		return err
	}

	resp.Consignment = req
	resp.Created = true
	return nil
}

func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	repo := h.getRepository()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		log.Printf("get consignments err: %v\n", err)
		return err
	}

	resp.Consignments = consignments
	return nil
}
