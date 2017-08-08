package usecase

import (
	"github.com/NgulikinApp/luffy/user"
	"github.com/NgulikinApp/luffy/user/repository"
)

type UserUsecase interface {
	GetByID(id int64) (*user.User, error)
	SignIn(email string, password string) (*user.User, error)
}

type userUsecase struct {
	UserRepository repository.UserRepository
}

func (self *userUsecase) GetByID(id int64) (*user.User, error) {
	return self.UserRepository.GetByID(id)
}

func (self *userUsecase) SignIn(email string, password string) (*user.User, error) {
	return self.UserRepository.SignIn(email, password)
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r}
}
