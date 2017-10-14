package usecase

import (
	"strconv"

	"github.com/NgulikinApp/luffy/product"
	"github.com/NgulikinApp/luffy/product/repository"
)

type Usecase interface {
	Store(p *product.Product) error
	Fetch(filter product.Filter, num int64, cursor int64) ([]*product.Product, string, error)
}

type usecase struct {
	Repo repository.Repository
}

func (u *usecase) Store(p *product.Product) error {
	errs := u.Repo.Store(p)
	return errs
}

func (u *usecase) Fetch(filter product.Filter, num int64, cursor int64) ([]*product.Product, string, error) {
	products, err := u.Repo.Fetch(filter, num, cursor)
	if err != nil {
		return nil, ``, err
	}

	if products == nil || len(products) == 0 {
		return make([]*product.Product, 0), ``, err
	}

	if len(products) > 0 {
		cursor = cursor + int64(len(products))
	}

	snc := strconv.FormatInt(cursor, 10)

	return products, snc, nil
}

func NewUsecase(r repository.Repository) Usecase {
	return &usecase{r}
}
