package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Aymenworks/go-practice/cache"
	"github.com/Aymenworks/go-practice/entity"
	"github.com/Aymenworks/go-practice/errors"
	"github.com/Aymenworks/go-practice/service"
)

var (
	postService service.PostService
	postCache   cache.PostCache
)

type controller struct{}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	postCache = cache
	return &controller{}
}

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	GetPostByID(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	posts, err := postService.GetAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error fetching post from DB"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) GetPostByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	postID := strings.Split(request.URL.Path, "/")[2]

	var post *entity.Post = postCache.Get(postID)
	if post == nil {
		log.Print("Not using cache")
		post, err := postService.FindByID(postID)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error fetching post from DB"})
			return
		}
		postCache.Set(postID, post)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	} else {
		log.Print("Using cache")
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	}
}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error marshalling the post request"})
		return
	}

	err1 := postService.Validate(&post)

	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	resultPost, err2 := postService.Create(&post)

	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post object"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resultPost)
}
