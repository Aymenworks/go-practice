package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Aymenworks/go-practiceclear
	/entity"
	"github.com/Aymenworks/go-practice/errors"
	"github.com/Aymenworks/go-practice/service"
)

var (
	postService service.PostService
)

type controller struct{}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
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
