//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	_ "github.com/google/wire"
	"go_online_course/internal/oauth/delivery"
	"go_online_course/internal/oauth/repository"
	"go_online_course/internal/oauth/usecase"
	repository2 "go_online_course/internal/user/repository"
	usecase2 "go_online_course/internal/user/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *oauthHandler.oauthHandler {
	wire.Build(delivery.NewOauthHandler,
		repository.NewOauthClientRepository,
		repository.NewOauthAccessTokenRepositor,
		repository.NewOauthRefreshTokenRepositoryRepositor,
		usecase.NewOauthUseCase,
		repository2.NewUserRepositoryImpl,
		usecase2.NewUserUseCase,
	)
}
