package repository

import (
	"context"
	"github.com/alexgtn/supernova/domain/user"
	"github.com/alexgtn/supernova/ent"
	"github.com/pkg/errors"
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

	dto, err := user.NewUser(u.ID, u.Age, u.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get user entity %d", id)
	}

	return dto, nil
}
