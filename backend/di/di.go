package di

import (
	"github.com/RyuichiroYoshida/quest-board-project/internal/auth/infrastructure"
	"github.com/RyuichiroYoshida/quest-board-project/internal/auth/interface/handler"
	"github.com/RyuichiroYoshida/quest-board-project/internal/auth/usecase"
	"gorm.io/gorm"
)

type Container struct {
	AuthHandler handler.AuthHandler
}

func InitContainer(db *gorm.DB) *Container {
	return &Container{
		AuthHandler: provideAuthHandler(db),
	}
}

// AuthHandlerを返すDIセットアップ関数
func provideAuthHandler(db *gorm.DB) handler.AuthHandler {
	authRepo := infrastructure.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	return handler.NewAuthHandler(*authUsecase)
}
