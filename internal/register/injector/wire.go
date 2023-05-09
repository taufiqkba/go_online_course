//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course/internal/register/delivery/http"
	"go_online_course/internal/register/usecase"
	"go_online_course/internal/user/repository"
	usecase2 "go_online_course/internal/user/usecase"
	"go_online_course/pkg/mail/sendgrid"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.RegisterHandler {
	wire.Build(
		http.NewRegisterHandler,
		usecase.NewRegisterUseCase,
		repository.NewUserRepositoryImpl,
		usecase2.NewUserUseCase,
		sendgrid.NewMail,
	)
	return &http.RegisterHandler{}
}
