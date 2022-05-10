package user

import (
	"errors"
	"time"
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

func (u *User) GetID() int {
	return u.id
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}
