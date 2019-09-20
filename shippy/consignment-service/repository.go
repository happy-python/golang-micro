package main

import (
	pb "golang-micro/shippy/consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	Db         = "shippy"
	Collection = "consignments"
)

type IRepository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type Repository struct {
	session *mgo.Session
}

func (repo *Repository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

// 获取全部数据
func (repo *Repository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

// 关闭连接
func (repo *Repository) Close() {
	// Close() 会在每次查询结束的时候关闭会话
	// Mgo 会在启动的时候生成一个 "主" 会话
	// 你可以使用 Copy() 直接从主会话复制出新会话来执行，即每个查询都会有自己的数据库会话
	// 同时每个会话都有自己连接到数据库的 socket 及错误处理，这么做既安全又高效
	// 如果只使用一个连接到数据库的主 socket 来执行查询，那很多请求处理都会阻塞
	// Mgo 因此能在不使用锁的情况下完美处理并发请求
	// 不过弊端就是，每次查询结束之后，必须确保数据库会话要手动 Close
	// 否则将建立过多无用的连接，白白浪费数据库资源
	repo.session.Close()
}

func (repo *Repository) collection() *mgo.Collection {
	return repo.session.DB(Db).C(Collection)
}
