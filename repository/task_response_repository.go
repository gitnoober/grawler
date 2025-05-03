package repository

import (
	"github.com/gitnoober/grawler/models"

	"gorm.io/gorm"
)

type TaskResponseRepository interface {
	CreateTaskResponse(taskResponse *models.TaskResponse) error
	GetTaskResponse(taskResponseID string) (*models.TaskResponse, error)
	GetAllTaskResponses() ([]*models.TaskResponse, error)
	GetUnprocessedResponses(limit int) ([]*models.TaskResponse, error)
	UpdateTaskResponseSummary(taskResponse *models.TaskResponse) error
}

type taskResponseRepository struct {
	db *gorm.DB
}

// GetTaskResponse implements TaskResponseRepository.
func (r *taskResponseRepository) GetTaskResponse(taskResponseID string) (*models.TaskResponse, error) {
	var taskResponse models.TaskResponse
	if err := r.db.Where("id = ?", taskResponseID).First(&taskResponse).Error; err != nil {
		return nil, err
	}
	return &taskResponse, nil
}

func (r *taskResponseRepository) CreateTaskResponse(taskResponse *models.TaskResponse) error {
	return r.db.Create(taskResponse).Error
}

func (r *taskResponseRepository) GetAllTaskResponses() ([]*models.TaskResponse, error) {
	var taskResponses []*models.TaskResponse
	if err := r.db.Find(&taskResponses).Error; err != nil {
		return nil, err
	}
	return taskResponses, nil
}

func (r *taskResponseRepository) GetUnprocessedResponses(limit int) ([]*models.TaskResponse, error) {
	var taskResponses []*models.TaskResponse
	if err := r.db.Where("summary_generated = ?", false).Limit(limit).Find(&taskResponses).Error; err != nil {
		return nil, err
	}
	return taskResponses, nil
}

func (r *taskResponseRepository) UpdateTaskResponseSummary(taskResponse *models.TaskResponse) error {
	return r.db.Model(taskResponse).Update("summary_generated", taskResponse.SummaryGenerated).Error
}
