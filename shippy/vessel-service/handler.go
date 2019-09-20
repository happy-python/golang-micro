package main

import (
	"context"
	"errors"
	pb "golang-micro/shippy/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"log"
)

type handler struct {
	session *mgo.Session
}

func (h *handler) getRepository() IRepository {
	// 从主会话中 Clone() 出新会话处理查询
	return &Repository{
		session: h.session.Clone(),
	}
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	repo := h.getRepository()
	defer repo.Close()

	if err := repo.Create(req); err != nil {
		log.Printf("create vessel err: %v\n", err)
		return err
	}

	resp.Vessel = req
	resp.Created = true
	return nil
}

func (h *handler) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	repo := h.getRepository()
	defer repo.Close()

	vessels, err := repo.GetAll()
	if err != nil {
		log.Printf("get all vessel err: %v\n", err)
		return err
	}

	for _, v := range vessels {
		if v.Capacity >= req.Capacity && v.MaxWeight >= req.MaxWeight {
			resp.Vessel = v
			return nil
		}
	}

	return errors.New("No vessel can't be use")
}
