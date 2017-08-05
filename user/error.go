package user

import "errors"

var (
	ErrAlreadyExists = errors.New(`User already exists`)
	ErrRequired      = errors.New(`Invalid data`)
	ErrNotFound      = errors.New(`ID Not found`)
	ErrIDParam       = errors.New(`ID must be integer`)
)
