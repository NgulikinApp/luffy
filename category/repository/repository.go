package repository

import "github.com/NgulikinApp/luffy/category"

type Repository interface {
	Fetch(num int64, cursor int64) ([]*category.Category, error)
}
