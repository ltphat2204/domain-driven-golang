package dto

import (
	"time"
)

type TaskCreateDTO struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	DueAt       *time.Time `json:"due_at"`
}

type TaskUpdateDTO struct {
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	Status      *string     `json:"status"`
	DueAt       *time.Time  `json:"due_at"`
}