package repository

import (
	"context"
	"github.com/ltphat2204/domain-driven-golang/category/domain"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Save(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	result := r.db.WithContext(ctx).Create(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (r *categoryRepository) FindByID(ctx context.Context, id uint) (*domain.Category, error) {
	var category domain.Category
	result := r.db.WithContext(ctx).First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *categoryRepository) FindCategories(ctx context.Context, query *domain.CategoryQuery) ([]*domain.Category, int, error) {
	db := r.db.WithContext(ctx).Model(&domain.Category{})

	if query.Search != "" {
		db = db.Where("name LIKE ? OR description LIKE ?", "%"+query.Search+"%", "%"+query.Search+"%")
	}

	var total int64
	db.Count(&total)

	var categories []*domain.Category
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

	if err := dbQuery.Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, int(total), nil
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	result := r.db.WithContext(ctx).Save(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (r *categoryRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&domain.Category{}, id)
	return result.Error
}