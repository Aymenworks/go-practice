package repository

import (
	"../entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	GetAll() ([]entity.Post, error)
}
