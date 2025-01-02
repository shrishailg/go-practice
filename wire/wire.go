//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/go-practice/internal/handlers"
	"github.com/go-practice/internal/repository"
	"github.com/go-practice/internal/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// func Initialize(db *gorm.DB) {
// 	// Initialize instances
// 	wire.Build(repository.NewRepository)
// 	wire.Build(service.NewUserService)
// 	wire.Build(handlers.NewUserHandler)
// 	wire.Build(handlers.NewRequestHandler)
// 	wire.Build(router.NewRouter)
// 	wire.Build(server.NewServer)

// 	// Create Event instance with Greeter
// }

func InitializeRepository(db *gorm.DB) *repository.Repository {
	wire.Build(repository.NewRepository)
	return nil
}

func InitializeUserService() *service.UserService {
	wire.Build(service.NewUserService)
	return nil
}

func InitializeUserHandler() *handlers.UserHandler {
	wire.Build(handlers.NewUserHandler)
	return nil
}

func InitializeRequestHandler() *handlers.RequestHandler {
	wire.Build(handlers.NewRequestHandler)
	return nil
}
