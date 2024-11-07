package constructor

import (
	"github.com/phn00dev/go-crud/internal/app"
	userConstructor "github.com/phn00dev/go-crud/internal/domain/user/constructor"

)

func InitDependencies(dependencies *app.Dependencies) {
	userConstructor.InitAUserAuthRequirements(dependencies.DB)
	userConstructor.InitUserRequirements(dependencies.DB)
}
