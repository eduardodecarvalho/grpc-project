package services

import "github.com/edaurdodecarvalho/grpc-projest"

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return &pb.User {
		Id: "1234",
		Name: req.GetName(),
		Email: req.GetEmail(),
	}, nil
}
