package usecase_test

import (
	"testing"

	"github.com/NgulikinApp/luffy/user"
	"github.com/NgulikinApp/luffy/user/repository/mocks"
	"github.com/NgulikinApp/luffy/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockUser = user.User{
		ID:        1,
		Username:  `my-username`,
		Fullname:  `my-fullname`,
		DOB:       `2017-01-01`,
		Gender:    `male`,
		Source:    `web`,
		Activated: true,
	}
)

// Test Usecase Get Company Data By ID
func TestUsecaseGetByID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockRepo.On("GetByID",
		mock.AnythingOfType("int64"),
	).Return(&mockUser, nil).Once()

	u := usecase.NewUserUsecase(mockRepo)
	res, err := u.GetByID(mockUser.ID)
	assert.NotNil(t, res)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

// Test Usecase Get Company Data By ID
func TestUsecaseSignin(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockRepo.On("SignIn",
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
	).Return(&mockUser, nil).Once()

	u := usecase.NewUserUsecase(mockRepo)
	res, err := u.SignIn(`email`, `password`)
	assert.NotNil(t, res)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
