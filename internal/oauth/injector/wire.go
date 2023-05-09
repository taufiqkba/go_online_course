//go:build wireinject
// +build wireinject

package injector

import (
	"go_online_course/internal/oauth/delivery"
	"go_online_course/internal/oauth/repository"
	"go_online_course/internal/oauth/usecase"
	repository2 "go_online_course/internal/user/repository"
	usecase2 "go_online_course/internal/user/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *delivery.OauthHandler {
	wire.Build(delivery.NewOauthHandler,
		repository.NewOauthClientRepository,
		repository.NewOauthAccessTokenRepository,
		repository.NewOauthRefreshTokenRepositoryRepository,
		usecase.NewOauthUseCase,
		repository2.NewUserRepositoryImpl,
		usecase2.NewUserUseCase,
	)
	return &delivery.OauthHandler{}
}
