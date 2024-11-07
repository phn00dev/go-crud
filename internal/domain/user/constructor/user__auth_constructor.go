package constructor

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-crud/internal/domain/user/handler"
	"github.com/phn00dev/go-crud/internal/domain/user/repository"
	"github.com/phn00dev/go-crud/internal/domain/user/service"
)

var (
	userAuthRepo    repository.AuthUserRepository
	userAuthService service.AuthUserService
	UserAuthHandler handler.AuthUserHandler
)

func InitAUserAuthRequirements(db *gorm.DB) {
	userAuthRepo = repository.NewAuthUserRepository(db)
	userAuthService = service.NewAuthUserService(userAuthRepo)
	UserAuthHandler = handler.NewAuthUserHandler(userAuthService)
}
