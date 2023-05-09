package repository

import (
	"go_online_course/internal/oauth/entity"
	"gorm.io/gorm"
)

type OauthAccessTokenRepository interface {
	Create(oauthAccessToken entity.OauthAccessToken) (*entity.OauthAccessToken, error)
	Delete(id int) error
}

type OauthAccessTokenImpl struct {
	db *gorm.DB
}

func (repository *OauthAccessTokenImpl) Create(oauthAccessToken entity.OauthAccessToken) (*entity.OauthAccessToken, error) {
	if result := repository.db.Create(oauthAccessToken).Error; result != nil {
		return nil, result
	}

	return &oauthAccessToken, nil
}

func (repository *OauthAccessTokenImpl) Delete(id int) error {
	var oauthAccessToken entity.OauthAccessToken

	if err := repository.db.Delete(&oauthAccessToken, id).Error; err != nil {
		return err
	}
	return nil
}

func NewOauthAccessTokenRepository(db *gorm.DB) OauthAccessTokenRepository {
	return &OauthAccessTokenImpl{db}
}
