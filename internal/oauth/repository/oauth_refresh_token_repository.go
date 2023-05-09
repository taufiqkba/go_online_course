package repository

import (
	"go_online_course/internal/oauth/entity"
	"gorm.io/gorm"
)

type OauthRefreshTokenRepository interface {
	Create(oauthRefreshToken entity.OauthRefreshToken) (*entity.OauthRefreshToken, error)
	FindOneByToken(token string) (*entity.OauthRefreshToken, error)
	Delete(id int) error
}

type OauthRefreshTokenRepositoryImpl struct {
	db *gorm.DB
}

func (repository *OauthRefreshTokenRepositoryImpl) Create(oauthRefreshToken entity.OauthRefreshToken) (*entity.OauthRefreshToken, error) {
	if err := repository.db.Create(oauthRefreshToken).Error; err != nil {
		return nil, err
	}
	return &oauthRefreshToken, nil
}

func (repository *OauthRefreshTokenRepositoryImpl) FindOneByToken(token string) (*entity.OauthRefreshToken, error) {
	//var oauthRefreshToken entity.OauthRefreshToken

	//if err := repository.db.Where("token = ?", token)
	// TODO Implement me
	panic("implement me!")
}

func (repository *OauthRefreshTokenRepositoryImpl) Delete(id int) error {
	var oauthRefreshToken entity.OauthRefreshToken

	if err := repository.db.Delete(&oauthRefreshToken, id).Error; err != nil {
		return err
	}
	return nil
}

func NewOauthRefreshTokenRepositoryRepository(db *gorm.DB) OauthRefreshTokenRepository {
	return &OauthRefreshTokenRepositoryImpl{db}
}
