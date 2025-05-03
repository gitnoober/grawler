package models

import "time"

type TaskResponse struct {
	ID uint `gorm:"primary_key"`
	TaskID uint `gorm:"not null"`
	UrlID uint `gorm:"not null"`
	Body string `gorm:"not null"`
	SummaryGenerated bool `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

