package grpc

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/alexgtn/supernova/domain/user"
	pb "github.com/alexgtn/supernova/proto"
)

type userService interface {
	GetByID(ctx context.Context, id int) (*user.User, error)
	Create(ctx context.Context, age int, name string) (*user.User, error)
	Update(ctx context.Context, id int, opts ...user.Option) (*user.User, error)
}

type delivery struct {
	pb.UnimplementedUserServiceServer
	userService userService
}

func NewUserDeliveryGrpc(r userService) *delivery {
	return &delivery{
		userService: r,
	}
}

func (d *delivery) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.OneUserReply, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, errors.Wrapf(err, "invalid user data %s", in.String()).Error())
	}

	u, err := d.userService.Create(ctx, int(in.Age), in.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrapf(err, "error creating user with data %s", in.String()).Error())
	}

	return &pb.OneUserReply{
		Id:        uint32(u.ID()),
		Age:       uint32(u.Age()),
		Name:      u.Name(),
		CreatedAt: timestamppb.New(u.GetCreatedAt()),
	}, nil
}

func (d *delivery) GetOne(ctx context.Context, in *pb.OneUserRequest) (*pb.OneUserReply, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, errors.Wrapf(err, "invalid user data %s", in.String()).Error())
	}

	u, err := d.userService.GetByID(ctx, int(in.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrapf(err, "error getting the user with id %d", in.GetId()).Error())
	}

	return &pb.OneUserReply{
		Id:        uint32(u.ID()),
		Age:       uint32(u.Age()),
		Name:      u.Name(),
		CreatedAt: timestamppb.New(u.GetCreatedAt()),
	}, nil
}

func (d *delivery) Update(ctx context.Context, in *pb.UpdateUserRequest) (*pb.OneUserReply, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, errors.Wrapf(err, "invalid user data %s", in.String()).Error())
	}

	u, err := d.userService.Update(ctx, int(in.GetId()), user.WithName(in.GetName()), user.WithAge(in.GetAge()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrapf(err, "error updating the user with id %d", in.GetId()).Error())
	}

	return &pb.OneUserReply{
		Id:        uint32(u.ID()),
		Age:       uint32(u.Age()),
		Name:      u.Name(),
		CreatedAt: timestamppb.New(u.GetCreatedAt()),
	}, nil
}
