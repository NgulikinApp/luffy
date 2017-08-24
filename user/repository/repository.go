package repository

import (
	"github.com/NgulikinApp/luffy/user"
)

type Repository interface {
	GetByID(id int64) (*user.User, error)
	Store(usr *user.User) error
}
