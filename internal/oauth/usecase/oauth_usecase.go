package usecase

import (
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go_online_course/internal/oauth/dto"
	"go_online_course/internal/oauth/entity"
	"go_online_course/internal/oauth/repository"
	"go_online_course/internal/user/usecase"
	"go_online_course/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type OauthUseCase interface {
	Login(loginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, error)
	Refresh(refreshTokenRequestBody dto.RefreshTokenRequestBody) (*dto.LoginResponse, error)
}

type OauthUseCaseImpl struct {
	oauthClientRepository       repository.OauthClientRepository
	oauthAccessTokenRepository  repository.OauthAccessTokenRepository
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository
	userUseCase                 usecase.UserUseCase
}

func (usecase *OauthUseCaseImpl) Login(loginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, error) {
	//	check oauth client_id and oauth_client_secret
	_, err := usecase.oauthClientRepository.FindByClientIdAndClientSecret(loginRequestBody.ClientID, loginRequestBody.ClientSecret)
	if err != nil {
		return nil, err
	}

	var user dto.UserResponse

	//	login using user data
	dataUser, err := usecase.userUseCase.FindByEmail(loginRequestBody.Email)
	if err != nil {
		return nil, errors.New("username or password is invalid")
	}

	user.ID = dataUser.ID
	user.Name = dataUser.Name
	user.Email = dataUser.Email
	user.Password = dataUser.Password

	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequestBody.Password))
	if err != nil {
		return nil, errors.New("username or password is invalid")
	}

	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	claims := &dto.ClaimResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	dataOauthAccessToken := entity.OauthAccessToken{
		OauthClient: &OauthClient.ID,
		UserID:      user.ID,
		Token:       tokenString,
		Scope:       "*",
		ExpiredAt: sql.NullTime{
			Time: expirationTime,
		},
	}

	oauthAccessToken, err := usecase.oauthAccessTokenRepository.Create(dataOauthAccessToken)

	if err != nil {
		return nil, err
	}

	dataOauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		UserID:             user.ID,
		Token:              utils.RandomString(128),
		ExpiredAt: sql.NullTime{
			Time: time.Now().Add(24 * 366 * time.Hour),
		},
	}

	oauthRefreshToken, err := usecase.oauthRefreshTokenRepository.Create(dataOauthRefreshToken)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: oauthRefreshToken,
		Type:         "",
		ExpiredAt:    "",
		Scope:        "",
	}, nil

}

func (usecase *OauthUseCaseImpl) Refresh(refreshTokenRequestBody dto.RefreshTokenRequestBody) (*dto.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewOauthUseCase() OauthUseCase {
	return &OauthUseCaseImpl{}
}
