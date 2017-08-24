package usecase

import (
	"github.com/NgulikinApp/luffy/user"
	"github.com/NgulikinApp/luffy/user/repository"
)

type Usecase interface {
	GetByID(id int64) (*user.User, error)
}

type usecase struct {
	repo repository.Repository
}

func (u *usecase) GetByID(id int64) (*user.User, error) {
	return u.repo.GetByID(id)
}

func NewUsecase(r repository.Repository) Usecase {
	return &usecase{r}
}
