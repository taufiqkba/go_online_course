package repository

import (
	"go_online_course/internal/oauth/entity"
	"gorm.io/gorm"
)

type OauthClientRepository interface {
	FindByClientIdAndClientSecret(clientId string, clientSecret string) (*entity.OauthClient, error)
}

type OauthClientImpl struct {
	db *gorm.DB
}

func (oc *OauthClientImpl) FindByClientIdAndClientSecret(clientId string, clientSecret string) (*entity.OauthClient, error) {
	var oauthClient entity.OauthClient

	if err := oc.db.Where("client_id = ?", clientId).Where("client_secret = ?", clientSecret).First(&oauthClient).Error; err != nil {
		return nil, err
	}

	return &oauthClient, nil
}

func NewOauthClientRepository() OauthClientRepository {
	return &OauthClientImpl{}
}
