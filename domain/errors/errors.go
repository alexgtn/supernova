package errors

import (
	"github.com/pkg/errors"
)

var (
	ErrNotFound    = errors.New("error not found")
	ErrNotSingular = errors.New("error not singular")
)
