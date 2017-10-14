package repository

import "github.com/NgulikinApp/luffy/product"

type Repository interface {
	Store(p *product.Product) error
	Fetch(filter product.Filter, num int64, cursor int64) ([]*product.Product, error)
}
