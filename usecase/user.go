package usecase

import (
	"context"
	"github.com/alexgtn/supernova/domain/user"
	pb "github.com/alexgtn/supernova/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userRepo interface {
	GetByID(ctx context.Context, id int) (*user.User, error)
	Create(ctx context.Context, age int, name string) (*user.User, error)
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

func (s *server) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.OneUserReply, error) {
	u, err := s.userRepo.Create(ctx, int(in.Age), in.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrapf(err, "error creating user with data %s", in.String()).Error())
	}

	return &pb.OneUserReply{
		Id:        uint32(u.GetID()),
		Age:       uint32(u.GetAge()),
		Name:      u.GetName(),
		CreatedAt: timestamppb.New(u.GetCreatedAt()),
	}, nil
}

func (s *server) GetOne(ctx context.Context, in *pb.OneUserRequest) (*pb.OneUserReply, error) {
	u, err := s.userRepo.GetByID(ctx, int(in.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrapf(err, "error getting the user with id %d", in.GetId()).Error())
	}

	return &pb.OneUserReply{
		Id:        uint32(u.GetID()),
		Age:       uint32(u.GetAge()),
		Name:      u.GetName(),
		CreatedAt: timestamppb.New(u.GetCreatedAt()),
	}, nil
}
