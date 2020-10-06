package cache

import (
	"github.com/Aymenworks/go-practice/entity"
)

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
