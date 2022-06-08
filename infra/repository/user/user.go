package user

import (
	"context"

	"github.com/pkg/errors"

	tx "github.com/alexgtn/supernova/common/db"
	errors2 "github.com/alexgtn/supernova/domain/errors"
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
		switch err.(type) {
		case *ent.NotFoundError:
			return nil, errors.Wrapf(errors2.ErrNotFound, "could not find user with id %d: %v", id, err)
		case *ent.NotSingularError:
			return nil, errors.Wrapf(errors2.ErrNotSingular, "more than one user with id %d: %v", id, err)
		default:
			return nil, errors.Wrapf(err, "could not get user by id %d", id)
		}
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

type Option func(u *ent.UserUpdateOne)

func WithName(name string) Option {
	return func(u *ent.UserUpdateOne) {
		u.SetName(name)
	}
}

func WithAge(age int) Option {
	return func(u *ent.UserUpdateOne) {
		u.SetAge(age)
	}
}

// OptionsFromDomain maps domain opts to repository opts
func OptionsFromDomain(opts ...user.Option) []Option {
	u := &user.User{}
	for _, opt := range opts {
		opt(u)
	}

	var outOpts []Option
	if u.Age() > 0 {
		outOpts = append(outOpts, WithAge(u.Age()))
	}

	if u.Name() != "" {
		outOpts = append(outOpts, WithName(u.Name()))
	}

	return outOpts
}

func (r *userRepo) Update(ctx context.Context, id int, opts ...Option) (*user.User, error) {
	client := tx.OrClient(ctx, r.client)
	uUpd := client.User.
		UpdateOneID(id)

	for _, optApplyTo := range opts {
		optApplyTo(uUpd) // apply option
	}

	u, err := uUpd.Save(ctx)
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
