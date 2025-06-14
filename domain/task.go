package domain

import (
	"context"
	"time"
)

type TaskStatus string

const (
	StatusPending TaskStatus = "Pending"
	StatusDoing   TaskStatus = "Doing"
	StatusDone    TaskStatus = "Done"
)

type Task struct {
	ID          uint        `gorm:"primaryKey"`
	Title       string      `gorm:"not null"`
	Description string
	Status      TaskStatus  `gorm:"type:varchar(10);default:'Pending'"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime"`
	DueAt       *time.Time  `gorm:"type:timestamp"`
}

type TaskRepository interface {
	Save(ctx context.Context, task *Task) (*Task, error)
	FindByID(ctx context.Context, id uint) (*Task, error)
	FindAll(ctx context.Context) ([]*Task, error)
	Update(ctx context.Context, task *Task) (*Task, error)
	Delete(ctx context.Context, id uint) error
}