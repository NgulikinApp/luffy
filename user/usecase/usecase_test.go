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
		ID:   1,
		Name: `User Name`,
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
