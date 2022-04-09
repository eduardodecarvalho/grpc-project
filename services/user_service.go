package services

import (
	"context"

	"github.com/eduardodecarvalho/grpc-project/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return &pb.User{
		Id:    "1234",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}
