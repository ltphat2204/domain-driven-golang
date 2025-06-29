package application

import (
	"context"
	"fmt"

	"github.com/ltphat2204/domain-driven-golang/modules/category/domain"
	"github.com/ltphat2204/domain-driven-golang/config"
	"github.com/ltphat2204/domain-driven-golang/utils"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, name, description string) (*domain.Category, error)
	GetCategoryByID(ctx context.Context, id uint) (*domain.Category, error)
	GetCategories(ctx context.Context, query *domain.CategoryQuery) ([]*domain.Category, int, error)
	UpdateCategory(ctx context.Context, id uint, name, description, color *string) (*domain.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
}

type categoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(ctx context.Context, name, description string) (*domain.Category, error) {
	color := utils.GetRandomColor(config.ColorPalette)
	category := &domain.Category{
		Name:        name,
		Description: description,
		Color:       color,
	}
	return s.repo.Save(ctx, category)
}

func (s *categoryService) GetCategoryByID(ctx context.Context, id uint) (*domain.Category, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *categoryService) GetCategories(ctx context.Context, query *domain.CategoryQuery) ([]*domain.Category, int, error) {
	return s.repo.FindCategories(ctx, query)
}

func (s *categoryService) UpdateCategory(ctx context.Context, id uint, name, description, color *string) (*domain.Category, error) {
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if name != nil && *name != "" {
		category.Name = *name
	}
	if description != nil {
		category.Description = *description
	}
	if color != nil && *color != "" {
		if !utils.IsValidColor(*color, config.ColorPalette) {
			return nil, fmt.Errorf("invalid color: must be one of %v", config.ColorPalette)
		}
		category.Color = *color
	}
	return s.repo.Update(ctx, category)
}

func (s *categoryService) DeleteCategory(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}