package repository

import (
	"github.com/gitnoober/grawler/models"

	"gorm.io/gorm"
)

type UrlRepository interface {
	CreateUrl(url *models.Url) error
	GetUrl(url string) (*models.Url, error)
}

type urlRepository struct {
	db *gorm.DB
}

func (r *urlRepository) CreateUrl(url *models.Url) error {
	return r.db.Create(url).Error
}

func (r *urlRepository) GetUrlMetadata(url string) (*models.Url, error) {
	var urlMetadata models.Url
	if err := r.db.Where("url = ?", url).First(&urlMetadata).Error; err != nil {
		return nil, err
	}
	return &urlMetadata, nil
}

