package main

import (
	"github.com/jinzhu/gorm"
	pb "golang-micro/shippy/user-service/proto/user"
)

type IRepository interface {
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	Create(*pb.User) error
	GetByEmail(string) (*pb.User, error)
}

type Repository struct {
	db *gorm.DB
}

func (repo *Repository) Get(id string) (*pb.User, error) {
	user := new(pb.User)
	if err := repo.db.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *Repository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *Repository) Create(user *pb.User) error {
	return repo.db.Create(&user).Error
}

func (repo *Repository) GetByEmail(email string) (*pb.User, error) {
	user := new(pb.User)
	if err := repo.db.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
