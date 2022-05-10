package usecase

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/alexgtn/supernova/common/db"
	"github.com/alexgtn/supernova/domain/user"
	user_repo "github.com/alexgtn/supernova/infra/repository/user"
	pb "github.com/alexgtn/supernova/proto"
)

type userRepo interface {
	GetByID(ctx context.Context, id int) (*user.User, error)
	Create(ctx context.Context, age int, name string) (*user.User, error)
	Update(ctx context.Context, id int, opts ...user_repo.Option) (*user.User, error)
}

type service struct {
	pb.UnimplementedUserServiceServer
	userRepo userRepo
}

func NewUserService(r userRepo) *service {
	return &service{
		userRepo: r,
	}
}

func (s *service) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.OneUserReply, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, errors.Wrapf(err, "invalid user data %s", in.String()).Error())
	}

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

func (s *service) GetOne(ctx context.Context, in *pb.OneUserRequest) (*pb.OneUserReply, error) {
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

func (s *service) Update(ctx context.Context, in *pb.UpdateUserRequest) (*pb.OneUserReply, error) {
	// start transaction
	txContext := db.NewTxContext(ctx)
	// override context
	ctx = txContext.GetContext()

	_, err := s.userRepo.Update(ctx, int(in.GetId()), user_repo.WithName(in.GetName()))
	if err != nil {
		return nil, txContext.ErrRollback(status.Errorf(codes.Internal, errors.Wrapf(err, "error updating the user with id %d", in.GetId()).Error()))
	}

	u, err := s.userRepo.Update(ctx, int(in.GetId()), user_repo.WithAge(int(in.GetAge())))
	if err != nil {
		return nil, txContext.ErrRollback(status.Errorf(codes.Internal, errors.Wrapf(err, "error updating the user with id %d", in.GetId()).Error()))
	}

	// commit transaction
	err = txContext.Commit()
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrapf(err, "error updating the user with id %d", in.GetId()).Error())
	}

	return &pb.OneUserReply{
		Id:        uint32(u.GetID()),
		Age:       uint32(u.GetAge()),
		Name:      u.GetName(),
		CreatedAt: timestamppb.New(u.GetCreatedAt()),
	}, nil
}
