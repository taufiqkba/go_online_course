package repository

import (
	"go_online_course/internal/oauth/entity"

	"gorm.io/gorm"
)

type OauthAccessTokenRepository interface {
	Create(oauthAccessToken entity.OauthAccessToken) (*entity.OauthAccessToken, error)
	Delete(id int) error
}

type OauthAccessTokenRepositoryImpl struct {
	//connect to database
	db *gorm.DB
}

// Create implements OauthAccessTokenRepository
func (oa *OauthAccessTokenRepositoryImpl) Create(oauthAccessToken entity.OauthAccessToken) (*entity.OauthAccessToken, error) {
	if result := oa.db.Create(oauthAccessToken).Error; result != nil {
		return nil, result
	}

	return &oauthAccessToken, nil
}

// Delete implements OauthAccessTokenRepository
func (oa *OauthAccessTokenRepositoryImpl) Delete(id int) error {
	var oauthAccessToken entity.OauthAccessToken
	if err := oa.db.Delete(&oauthAccessToken, id).Error; err != nil {
		return err
	}
	return nil
}

func NewOauthAccessTokenRepository() OauthAccessTokenRepository {
	return &OauthAccessTokenRepositoryImpl{}
}
