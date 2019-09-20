package main

import (
	pb "golang-micro/shippy/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

const (
	Db         = "shippy"
	Collection = "vessels"
)

type IRepository interface {
	Create(*pb.Vessel) error
	GetAll() ([]*pb.Vessel, error)
	Close()
}

type Repository struct {
	session *mgo.Session
}

func (repo *Repository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// 获取全部数据
func (repo *Repository) GetAll() ([]*pb.Vessel, error) {
	var vessels []*pb.Vessel
	err := repo.collection().Find(nil).All(&vessels)
	return vessels, err
}

func (repo *Repository) Close() {
	repo.session.Close()
}

func (repo *Repository) collection() *mgo.Collection {
	return repo.session.DB(Db).C(Collection)
}
