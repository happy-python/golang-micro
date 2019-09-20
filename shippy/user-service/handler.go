package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	pb "golang-micro/shippy/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const topic = "user.created"

type handler struct {
	repo IRepository
}

func (h *handler) getRepository() IRepository {
	return &Repository{}
}

func (h *handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	// 哈希处理用户输入的密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPwd)
	if err = h.repo.Create(req); err != nil {
		log.Printf("create user err: %v\n", err)
		return err
	}

	// 创建用户成功，发送带有用户信息的消息
	if err = micro.NewPublisher(topic, client.DefaultClient).Publish(ctx, req); err != nil {
		log.Printf("publish err: %v\n", err)
		return err
	}

	resp.User = req
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	user, err := h.repo.Get(req.Id)
	if err != nil {
		log.Printf("get user err: %v\n", err)
		return err
	}

	resp.User = user
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		log.Printf("get all user err: %v\n", err)
		return err
	}

	resp.Users = users
	return nil
}

// 生成token
func (h *handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	user, err := h.repo.GetByEmail(req.Email)
	if err != nil {
		log.Printf("get user by email err: %v\n", err)
		return err
	}

	// 进行密码验证
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Printf("compare password err: %v\n", err)
		return err
	}

	service := TokenService{}
	token, err := service.Encode(user)
	if err != nil {
		log.Printf("encode err: %v\n", err)
		return err
	}

	resp.Token = token
	return nil
}

// 校验token
func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	service := TokenService{}
	claims, err := service.Decode(req.Token)
	if err != nil {
		log.Printf("decode err: %v\n", err)
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	resp.Valid = true
	return nil
}
