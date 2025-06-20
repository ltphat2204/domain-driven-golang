package infrastructure

import (
	"context"
	"github.com/ltphat2204/domain-driven-golang/task/domain"
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
	result := r.db.WithContext(ctx).Preload("Category").First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r *taskRepository) FindTasks(ctx context.Context, query *domain.TaskQuery) ([]*domain.Task, int, error) {
	db := r.db.WithContext(ctx).Preload("Category").Model(&domain.Task{})

	if query.Search != "" {
		db = db.Where("title LIKE ? OR description LIKE ?", "%"+query.Search+"%", "%"+query.Search+"%")
	}

	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}

	var total int64
	db.Count(&total)

	var tasks []*domain.Task
	dbQuery := db
	if query.SortBy != "" && query.SortOrder != "" {
		dbQuery = db.Order(query.SortBy + " " + query.SortOrder)
	} else {
		dbQuery = db.Order("created_at desc")
	}

	if query.Page > 0 && query.PageSize > 0 {
		offset := (query.Page - 1) * query.PageSize
		dbQuery = dbQuery.Offset(offset).Limit(query.PageSize)
	}

	if err := dbQuery.Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, int(total), nil
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