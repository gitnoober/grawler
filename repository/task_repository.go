package repository

import (
	"fmt"

	"github.com/gitnoober/grawler/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTask(taskID string) (*models.Task, error)
	UpdateTask(task *models.Task) error
}

type taskRepository struct {
	db *gorm.DB
}


func (r *taskRepository) CreateTask(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetTask(taskID string) (*models.Task, error) {
	var task models.Task
	if err := r.db.Where("id = ?", taskID).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) UpdateTask(task *models.Task) error {
	fmt.Println("Updating task: ", task.ID)
	fmt.Println("Task: ", task)
	if err := r.db.Model(&task).Where("id = ?", task.ID).Updates(task).Error; err != nil {
		fmt.Println("Error updating task: ", err)
		return err
	}
	return nil
}