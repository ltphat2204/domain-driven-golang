package domain

import (
	"context"
	"time"
	"github.com/ltphat2204/domain-driven-golang/common"
)

type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	Color       string    `gorm:"type:varchar(7)"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

type CategoryQuery struct {
	common.BaseQuery
	Search    string
	SortBy    string
	SortOrder string
}

type CategoryRepository interface {
	Save(ctx context.Context, category *Category) (*Category, error)
	FindByID(ctx context.Context, id uint) (*Category, error)
	FindCategories(ctx context.Context, query *CategoryQuery) ([]*Category, int, error)
	Update(ctx context.Context, category *Category) (*Category, error)
	Delete(ctx context.Context, id uint) error
}