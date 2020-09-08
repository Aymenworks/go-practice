package service

import (
	"testing"

	"../entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) GetAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestGetAll(t *testing.T) {
	mockRepo := new(MockRepository)

	mockPost := entity.Post{ID: 1, Title: "A", Text: "B"}

	// Setup expectations
	mockRepo.On("GetAll").Return([]entity.Post{mockPost}, nil)
	assert.NotNil(t, mockRepo)

	testService := NewPostService(mockRepo)
	result, _ := testService.GetAll()

	mockRepo.AssertExpectations(t)
	var identifier int64 = 1
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	mockPost := entity.Post{ID: 1, Title: "A", Text: "B"}

	// Setup expectations
	mockRepo.On("Save").Return(&mockPost, nil)
	assert.NotNil(t, mockRepo)

	testService := NewPostService(mockRepo)
	result, _ := testService.Create(&mockPost)

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "Post should not be nil")
}

func TestValidateEmptyTitle(t *testing.T) {
	post := entity.Post{
		ID:    1,
		Title: "",
		Text:  "Empty text",
	}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "Post.Title should not be empty")
}

func TestValidateSuccess(t *testing.T) {
	post := entity.Post{
		ID:    1,
		Title: "Some title",
		Text:  "Empty text",
	}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.Nil(t, err)
}
