package dto

type TaskCreateDTO struct {
    Title       string `json:"title" binding:"required"`
    Description string `json:"description"`
}

type TaskUpdateDTO struct {
    Title       *string `json:"title"`
    Description *string `json:"description"`
    Status      *string `json:"status"`
}