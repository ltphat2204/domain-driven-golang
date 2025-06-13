package domain

import (
	"context"
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	Description string
	Status      string    `gorm:"default:'pending'"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

type TaskRepository interface {
	Save(ctx context.Context, task *Task) (*Task, error)
	FindByID(ctx context.Context, id uint) (*Task, error)
	FindAll(ctx context.Context) ([]*Task, error)
	Update(ctx context.Context, task *Task) (*Task, error)
	Delete(ctx context.Context, id uint) error
}