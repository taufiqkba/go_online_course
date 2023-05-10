package usecase

import (
	"database/sql"
	"errors"
	usecase2 "go_online_course/internal/admin/usecase"
	"go_online_course/internal/oauth/dto"
	"go_online_course/internal/oauth/entity"
	"go_online_course/internal/oauth/repository"
	"go_online_course/internal/user/usecase"
	"go_online_course/pkg/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
	adminUseCase                usecase2.AdminUseCase
}

// Login implements OauthUseCase
func (usecase *OauthUseCaseImpl) Login(loginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, error) {
	// check oauth client_id and oauth_client_secret
	oauthClient, err := usecase.oauthClientRepository.FindByClientIdAndClientSecret(loginRequestBody.ClientID, loginRequestBody.ClientSecret)
	if err != nil {
		return nil, err
	}

	var user dto.UserResponse

	///CHECK LOGIN ADMIN
	if oauthClient.Name == "web-admin" {
		dataAdmin, err := usecase.adminUseCase.FindByEmail(loginRequestBody.Email)
		if err != nil {
			return nil, errors.New("email or password is invalid")
		}
		user.ID = dataAdmin.ID
		user.Email = dataAdmin.Email
		user.Password = dataAdmin.Password

	} else {
		//CHECK LOGIN USER
		dataUser, err := usecase.userUseCase.FindByEmail(loginRequestBody.Email)

		if err != nil {
			return nil, errors.New("username or password is invalid")
		}

		user.ID = dataUser.ID
		user.Name = dataUser.Name
		user.Email = dataUser.Email
		user.Password = dataUser.Password
	}

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	// Compare login password valid or not
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequestBody.Password))
	if err != nil {
		return nil, errors.New("username or password is invalid")
	}

	/// CREATE JWT TOKEN STEPS
	// create Expiration of JWT Token
	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	claims := &dto.ClaimResponse{
		ID:      user.ID,
		Email:   user.Email,
		Name:    user.Name,
		IsAdmin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	if oauthClient.Name != "web-admin" {
		claims.IsAdmin = false
	}

	// create JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	// insert data to table oauth_access_token
	dataOauthAccessToken := entity.OauthAccessToken{
		OauthClientID: &oauthClient.ID,
		UserID:        user.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt: sql.NullTime{
			Time: expirationTime,
		},
	}

	oauthAccessToken, err := usecase.oauthAccessTokenRepository.Create(dataOauthAccessToken)

	if err != nil {
		return nil, err
	}

	// insert data to oauth_refresh_token table
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

	// return response
	return &dto.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: oauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil

}

// Refresh implements OauthUseCase
func (usecase *OauthUseCaseImpl) Refresh(refreshTokenRequestBody dto.RefreshTokenRequestBody) (*dto.LoginResponse, error) {
	panic("unimplemented")
}

func NewOauthUseCase(
	oauthClientRepository repository.OauthClientRepository,
	oauthAccessTokenRepository repository.OauthAccessTokenRepository,
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository,
	userUseCase usecase.UserUseCase,
	adminUseCase usecase2.AdminUseCase,
) OauthUseCase {
	return &OauthUseCaseImpl{
		oauthClientRepository,
		oauthAccessTokenRepository,
		oauthRefreshTokenRepository,
		userUseCase,
		adminUseCase,
	}
}
