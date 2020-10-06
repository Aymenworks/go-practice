package repository

import (
	"github.com/Aymenworks/go-practice/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	GetAll() ([]entity.Post, error)
	FindByID(id string) (*entity.Post, error)
}
