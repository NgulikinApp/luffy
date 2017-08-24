package user

import "errors"

var (
	ErrNotFound = errors.New(`ID Not found`)
	ErrIDParam  = errors.New(`ID must be integer`)
)

type ErrRequired struct {
	Message string
}

func (e ErrRequired) Error() string {
	return e.Message
}

type ErrAlreadyExists struct {
	Message string
}

func (e ErrAlreadyExists) Error() string {
	return e.Message
}
