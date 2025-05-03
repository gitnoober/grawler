package models

import "time"

type Task struct {
	ID uint `gorm:"primary_key"`
	Status string `gorm:"not null"`
	UrlID uint `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

const (
	TaskStatusPending = "pending"
	TaskStatusRunning = "running"
	TaskStatusCompleted = "completed"
	TaskStatusFailed = "failed"
)