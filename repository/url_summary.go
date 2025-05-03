package repository

import (
	"github.com/gitnoober/grawler/models"
	"gorm.io/gorm"
)

type UrlSummaryRepository interface {
	CreateUrlSummary(urlSummary *models.UrlSummary) error
	GetUrlSummary(urlSummaryID uint) (*models.UrlSummary, error)
	GetAllUrlSummaries() ([]*models.UrlSummary, error)
	UpdateUrlSummary(urlSummary *models.UrlSummary) error
}

type urlSummaryRepository struct {
	db *gorm.DB
}

func (r *urlSummaryRepository) CreateUrlSummary(urlSummary *models.UrlSummary) error {
	return r.db.Create(urlSummary).Error
}

func (r *urlSummaryRepository) GetUrlSummary(urlSummaryID uint) (*models.UrlSummary, error) {
	var urlSummary models.UrlSummary
	if err := r.db.First(&urlSummary, urlSummaryID).Error; err != nil {
		return nil, err
	}
	return &urlSummary, nil
}

func (r *urlSummaryRepository) GetAllUrlSummaries() ([]*models.UrlSummary, error) {
	var urlSummaries []*models.UrlSummary
	if err := r.db.Find(&urlSummaries).Error; err != nil {
		return nil, err
	}
	return urlSummaries, nil
}

func (r *urlSummaryRepository) UpdateUrlSummary(urlSummary *models.UrlSummary) error {
	return r.db.Save(urlSummary).Error
}



