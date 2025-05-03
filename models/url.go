package models

import "time"

type Url struct {
	ID  uint `gorm:"primary_key"`
	Url string `gorm:"not null"`
	MaxDepth int `gorm:"not null"`
	IsDisabled bool `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
