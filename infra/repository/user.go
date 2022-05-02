package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alexgtn/supernova/domain/user"
	"github.com/alexgtn/supernova/ent"
)

type userRepo struct {
	client *ent.Client
}

func NewUser(c *ent.Client) *userRepo {
	return &userRepo{c}
}

func (r *userRepo) GetByID(ctx context.Context, id int) (*user.User, error) {
	u, err := r.client.User.Get(ctx, id)
	if err != nil {
		_, ok := err.(*ent.NotFoundError)
		if ok {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "error fetching user %d", id)
	}

	dto, err := user.NewUser(u.ID, u.Age, u.Name, u.CreatedAt)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get user entity %d", id)
	}

	return dto, nil
}

func (r *userRepo) Create(ctx context.Context, age int, name string) (*user.User, error) {
	u, err := r.client.User.
		Create().
		SetName(name).
		SetAge(age).
		Save(ctx)
	if err != nil {
		_, ok := err.(*ent.NotFoundError)
		if ok {
			return nil, nil
		}
		return nil, errors.Wrap(err, "error creating user")
	}

	dto, err := user.NewUser(u.ID, u.Age, u.Name, u.CreatedAt)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get user entity %d", u.ID)
	}

	return dto, nil
}
