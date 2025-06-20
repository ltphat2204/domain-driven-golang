package dto

import (
	"time"
	"github.com/ltphat2204/domain-driven-golang/task/domain"
	"github.com/ltphat2204/domain-driven-golang/common"
)

type TaskCreateDTO struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	DueAt       *time.Time `json:"due_at"`
	CategoryID  *uint      `json:"category_id"`
}

type TaskUpdateDTO struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Status      *string    `json:"status"`
	DueAt       *time.Time `json:"due_at"`
	CategoryID  *uint      `json:"category_id"`
}

type TaskQueryDTO struct {
	Page      int    `form:"page" binding:"omitempty,gte=1"`
	PageSize  int    `form:"page_size" binding:"omitempty,gte=1"`
	Search    string `form:"search"`
	SortBy    string `form:"sort_by"`
	SortOrder string `form:"sort_order"`
	Status    string `form:"status"`
}

type TaskListResponse struct {
	Tasks []*domain.Task      `json:"tasks"`
	Meta  common.PaginationMeta `json:"meta"`
}