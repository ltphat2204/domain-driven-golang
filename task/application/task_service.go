package application

import (
	"context"
	"time"
	"github.com/ltphat2204/domain-driven-golang/task/domain"
)

type TaskService interface {
	CreateTask(ctx context.Context, title, description string, dueAt *time.Time, categoryID *uint) (*domain.Task, error)
	GetTaskByID(ctx context.Context, id uint) (*domain.Task, error)
	GetTasks(ctx context.Context, query *domain.TaskQuery) ([]*domain.Task, int, error)
	UpdateTask(ctx context.Context, id uint, title, description *string, status *domain.TaskStatus, dueAt *time.Time, categoryID *uint) (*domain.Task, error)
	DeleteTask(ctx context.Context, id uint) error
}

type taskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context, title, description string, dueAt *time.Time, categoryID *uint) (*domain.Task, error) {
	task := &domain.Task{
		Title:       title,
		Description: description,
		Status:      domain.StatusPending,
		DueAt:       dueAt,
		CategoryID:  categoryID,
	}
	return s.repo.Save(ctx, task)
}

func (s *taskService) GetTaskByID(ctx context.Context, id uint) (*domain.Task, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *taskService) GetTasks(ctx context.Context, query *domain.TaskQuery) ([]*domain.Task, int, error) {
	return s.repo.FindTasks(ctx, query)
}

func (s *taskService) UpdateTask(ctx context.Context, id uint, title, description *string, status *domain.TaskStatus, dueAt *time.Time, categoryID *uint) (*domain.Task, error) {
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if title != nil && *title != "" {
		task.Title = *title
	}
	if description != nil {
		task.Description = *description
	}
	if status != nil {
		task.Status = *status
	}
	if dueAt != nil {
		task.DueAt = dueAt
	}
	task.CategoryID = categoryID // Allow null to remove category
	return s.repo.Update(ctx, task)
}

func (s *taskService) DeleteTask(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}