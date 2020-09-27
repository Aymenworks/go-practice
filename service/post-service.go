package service

import (
	"errors"
	"math/rand"

	"github.com/Aymenworks/go-practice/entity"
	"github.com/Aymenworks/go-practice/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	GetAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("Post should not be nil")
	}
	if post.Title == "" {
		return errors.New("Post.Title should not be empty")
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) GetAll() ([]entity.Post, error) {
	return repo.GetAll()
}
