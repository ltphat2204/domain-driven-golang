package dto

import (
	"github.com/ltphat2204/domain-driven-golang/category/domain"
	"github.com/ltphat2204/domain-driven-golang/common"
)

type CategoryCreateDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CategoryUpdateDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
}

type CategoryQueryDTO struct {
	Page      int    `form:"page" binding:"omitempty,gte=1"`
	PageSize  int    `form:"page_size" binding:"omitempty,gte=1"`
	Search    string `form:"search"`
	SortBy    string `form:"sort_by"`
	SortOrder string `form:"sort_order"`
}

type CategoryListResponse struct {
	Categories []*domain.Category    `json:"categories"`
	Meta       common.PaginationMeta `json:"meta"`
}