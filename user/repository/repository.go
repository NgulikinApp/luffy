package repository

import (
	"github.com/NgulikinApp/luffy/user"
)

type UserRepository interface {
	GetByID(id int64) (*user.User, error)
}
