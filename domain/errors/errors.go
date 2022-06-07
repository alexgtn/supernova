package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type ErrNotFound struct {
	err error
}

func (e *ErrNotFound) Error() string {
	return e.err.Error()
}

func NewErrNotFound(err string) error {
	return &ErrNotFound{err: fmt.Errorf("error not found: %s", err)}
}

func WrapErrNotFound(err error) error {
	return &ErrNotFound{err: errors.Wrapf(err, "error not found: %s", err)}
}

type ErrNotSingular struct {
	err error
}

func (e *ErrNotSingular) Error() string {
	return e.err.Error()
}

func NewErrNotSingular(err string) error {
	return &ErrNotSingular{err: fmt.Errorf("error not singular: %s", err)}
}

func WrapErrNotSingular(err error) error {
	return &ErrNotSingular{err: errors.Wrapf(err, "error not singular: %s", err)}
}
