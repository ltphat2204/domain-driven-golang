package domain

import (
	"context"
	"time"

	"github.com/ltphat2204/domain-driven-golang/modules/category/domain"
	"github.com/ltphat2204/domain-driven-golang/common"
)

type TaskStatus string

const (
	StatusPending TaskStatus = "Pending"
	StatusDoing   TaskStatus = "Doing"
	StatusDone    TaskStatus = "Done"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Status      TaskStatus       `gorm:"type:varchar(10);default:'Pending'"`
	CreatedAt   time.Time        `gorm:"autoCreateTime"`
	UpdatedAt   time.Time        `gorm:"autoUpdateTime"`
	DueAt       *time.Time       `gorm:"type:timestamp"`
	CategoryID  *uint            `gorm:"foreignKey:CategoryID"` // Foreign key for Category
	Category    *domain.Category `gorm:"foreignKey:CategoryID"` // Association with Category
}

type TaskQuery struct {
	common.BaseQuery
	Search    string
	SortBy    string
	SortOrder string
	Status    *TaskStatus
}

type TaskRepository interface {
	Save(ctx context.Context, task *Task) (*Task, error)
	FindByID(ctx context.Context, id uint) (*Task, error)
	FindTasks(ctx context.Context, query *TaskQuery) ([]*Task, int, error)
	Update(ctx context.Context, task *Task) (*Task, error)
	Delete(ctx context.Context, id uint) error
}

func IsValidTaskStatus(status TaskStatus) bool {
	switch status {
	case StatusPending, StatusDoing, StatusDone:
		return true
	default:
		return false
	}
}