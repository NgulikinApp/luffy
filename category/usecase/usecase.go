package usecase

import (
	"github.com/NgulikinApp/luffy/category"
	"github.com/NgulikinApp/luffy/category/repository"
)

type Usecase interface {
	Fetch(num int64, cursor int64) ([]*category.Category, error)
}

type usecase struct {
	repo repository.Repository
}

func (u *usecase) Fetch(num int64, cursor int64) ([]*category.Category, error) {
	return u.repo.Fetch(num, cursor)
}

func NewUsecase(r repository.Repository) Usecase {
	return &usecase{r}
}
