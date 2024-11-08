package constructor

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-crud/internal/domain/post/handler"
	postRepository "github.com/phn00dev/go-crud/internal/domain/post/repository"
	"github.com/phn00dev/go-crud/internal/domain/post/service"
	"github.com/phn00dev/go-crud/internal/domain/user/repository"
	"github.com/phn00dev/go-crud/pkg/config"

)

var (
	postRepo    postRepository.PostRepository
	userRepo    repository.UserRepository
	postService service.PostService
	PostHandler handler.PostHandler
)

func InitPostRequirements(db *gorm.DB, config *config.Config) {
	postRepo = postRepository.NewPostRepository(db)
	userRepo = repository.NewUserRepository(db)
	postService = service.NewPostService(postRepo, userRepo, config)
	PostHandler = handler.NewPostHandler(postService)
}

