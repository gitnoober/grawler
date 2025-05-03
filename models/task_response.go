package models

import "time"

type TaskResponse struct {
	ID uint `gorm:"primary_key"`
	Body string `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

