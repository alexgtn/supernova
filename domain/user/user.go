package user

import (
	"errors"
)

type User struct {
	id   int
	age  int
	name string
}

func NewUser(id int, age int, name string) (*User, error) {
	if age <= 0 {
		return nil, errors.New("age should be a positive number")
	}
	if name == "" {
		return nil, errors.New("name is empty")
	}

	return &User{id, age, name}, nil
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
