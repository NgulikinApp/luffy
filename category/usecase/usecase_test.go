package usecase_test

import (
	"testing"

	"github.com/NgulikinApp/luffy/category"
	"github.com/NgulikinApp/luffy/category/repository/mocks"
	"github.com/NgulikinApp/luffy/category/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockCategory = category.Category{
		ID:           1,
		Name:         `technology`,
		ThumbnailURL: `http://thumbnail.com`,
	}
)

func TestUsecaseFetch(t *testing.T) {
	mockCategories := make([]*category.Category, 0)
	mockCategories = append(mockCategories, &mockCategory)
	mockRepo := new(mocks.Repository)

	mockRepo.On("Fetch",
		mock.AnythingOfType("int64"),
		mock.AnythingOfType("int64"),
	).Return(mockCategories, nil).Once()

	u := usecase.NewUsecase(mockRepo)
	res, err := u.Fetch(int64(10), int64(0))
	assert.NotNil(t, res)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
