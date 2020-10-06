package main

import (
	"os"

	"github.com/Aymenworks/go-practice/cache"
	"github.com/Aymenworks/go-practice/controller"
	router "github.com/Aymenworks/go-practice/http"
	"github.com/Aymenworks/go-practice/repository"
	"github.com/Aymenworks/go-practice/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestorePostRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	postController controller.PostController = controller.NewPostController(postService, postCache)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(os.Getenv("GO_PRACTICE_PORT"))
}
