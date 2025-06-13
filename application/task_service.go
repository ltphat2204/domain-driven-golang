package application

import (
	"context"
	"github.com/ltphat2204/domain-driven-golang/domain"
)

type TaskService interface {
	CreateTask(ctx context.Context, title, description string) (*domain.Task, error)
	GetTaskByID(ctx context.Context, id uint) (*domain.Task, error)
	GetAllTasks(ctx context.Context) ([]*domain.Task, error)
	UpdateTask(ctx context.Context, id uint, title, description, status *string) (*domain.Task, error)
	DeleteTask(ctx context.Context, id uint) error
}

type taskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context, title, description string) (*domain.Task, error) {
	task := &domain.Task{
		Title:       title,
		Description: description,
		Status:      "pending",
	}
	return s.repo.Save(ctx, task)
}

func (s *taskService) GetTaskByID(ctx context.Context, id uint) (*domain.Task, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *taskService) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	return s.repo.FindAll(ctx)
}

func (s *taskService) UpdateTask(ctx context.Context, id uint, title, description, status *string) (*domain.Task, error) {
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
	if status != nil && *status != "" {
		task.Status = *status
	}
	return s.repo.Update(ctx, task)
}

func (s *taskService) DeleteTask(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}