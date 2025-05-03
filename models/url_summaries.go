package models

import (
	"time"
)

type UrlSummary struct {
	ID        uint           `gorm:"primaryKey"`
	UrlID     uint           `gorm:"not null"`
	TaskID    uint           `gorm:"unique;not null"`
	Summary   string         `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}