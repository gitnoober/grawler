package repository

import (
	"gorm.io/gorm"
)


func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{db: db}
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}