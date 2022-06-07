package usecase

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (s *service) Create(ctx context.Context, age int, name string) (*user.User, error) {
	u, err := s.userRepo.Create(ctx, age, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrap(err, "error creating user").Error())
	}

	return u, nil
}

func (s *service) GetByID(ctx context.Context, id int) (*user.User, error) {
	u, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrapf(err, "error getting the user with id %d", id).Error())
	}

	return u, nil
}

func (s *service) Update(ctx context.Context, id int, opts ...user.Option) (*user.User, error) {
	// start transaction
	txContext := db.NewTxContext(ctx)
	// override context
	ctx = txContext.GetContext()

	u, err := s.userRepo.Update(ctx, id, user_repo.OptionsFromDomain(opts...)...)
	if err != nil {
		return nil, txContext.ErrRollback(status.Errorf(codes.Internal, errors.Wrapf(err, "error updating the user with id %d", id).Error()))
	}

	// commit transaction
	err = txContext.Commit()
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrapf(err, "error updating the user with id %d", id).Error())
	}

	return u, nil
}
