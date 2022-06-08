package user

import (
	"errors"
	"time"

	"golang.org/x/exp/constraints"
)

type User struct {
	id        int
	age       int
	name      string
	createdAt time.Time
}

func NewUser(id, age int, name string, createdAt time.Time) (*User, error) {
	if age <= 0 {
		return nil, errors.New("age should be a positive number")
	}

	if name == "" {
		return nil, errors.New("name is empty")
	}

	if createdAt.IsZero() {
		return nil, errors.New("createdAt invalid")
	}

	return &User{id, age, name, createdAt}, nil
}

func (u *User) ID() int {
	return u.id
}

func (u *User) Age() int {
	return u.age
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

type Option func(u *User)

func WithName(name string) Option {
	return func(u *User) {
		u.SetName(name)
	}
}

func WithAge[T constraints.Integer](age T) Option {
	return func(u *User) {
		u.SetAge(int(age))
	}
}
