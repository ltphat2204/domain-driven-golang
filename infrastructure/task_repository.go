package infrastructure

import (
	"context"
	"github.com/ltphat2204/domain-driven-golang/domain"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Save(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	result := r.db.WithContext(ctx).Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (r *taskRepository) FindByID(ctx context.Context, id uint) (*domain.Task, error) {
	var task domain.Task
	result := r.db.WithContext(ctx).First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r *taskRepository) FindAll(ctx context.Context) ([]*domain.Task, error) {
	var tasks []*domain.Task
	result := r.db.WithContext(ctx).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *taskRepository) Update(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	result := r.db.WithContext(ctx).Save(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (r *taskRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&domain.Task{}, id)
	return result.Error
}