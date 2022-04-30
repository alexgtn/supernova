package usecase

import (
	"context"
	"github.com/alexgtn/supernova/domain/user"
	pb "github.com/alexgtn/supernova/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userRepo interface {
	GetByID(ctx context.Context, id int) (*user.User, error)
}

type server struct {
	pb.UnimplementedUserServiceServer
	userRepo userRepo
}

func NewUserService(r userRepo) *server {
	return &server{
		userRepo: r,
	}
}

func (s *server) GetOne(ctx context.Context, in *pb.OneUserRequest) (*pb.OneUserReply, error) {
	u, err := s.userRepo.GetByID(ctx, int(in.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			errors.Wrapf(err, "error getting the user with id %d", in.GetId()).Error())
	}

	return &pb.OneUserReply{
		Id:   uint32(u.GetID()),
		Age:  uint32(u.GetAge()),
		Name: u.GetName(),
	}, nil
}
