package handlers

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/domain-driven-golang/task/application"
	"github.com/ltphat2204/domain-driven-golang/common"
	"github.com/ltphat2204/domain-driven-golang/task/domain"
	"github.com/ltphat2204/domain-driven-golang/task/dto"
)

type TaskHandler struct {
	service application.TaskService
}

func NewTaskHandler(service application.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input dto.TaskCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
		return
	}

	task, err := h.service.CreateTask(c.Request.Context(), input.Title, input.Description, input.DueAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Failed to create task", err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(task))
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid ID"))
		return
	}

	task, err := h.service.GetTaskByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewErrorResponse(http.StatusNotFound, "Task not found", err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(task))
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	var queryDTO dto.TaskQueryDTO
	if err := c.ShouldBindQuery(&queryDTO); err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
		return
	}

	// Set defaults
	page := 1
	if queryDTO.Page > 0 {
		page = queryDTO.Page
	}
	pageSize := 10
	if queryDTO.PageSize > 0 {
		pageSize = queryDTO.PageSize
	}

	allowedSortFields := []string{"title", "due_at", "created_at"}
	allowedSortOrders := []string{"asc", "desc"}

	if queryDTO.SortBy != "" && !slices.Contains(allowedSortFields, queryDTO.SortBy) {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid sort_by"))
		return
	}

	if queryDTO.SortOrder != "" && !slices.Contains(allowedSortOrders, queryDTO.SortOrder) {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid sort_order"))
		return
	}

	query := &domain.TaskQuery{
		BaseQuery: common.BaseQuery{
			Page:      page,
			PageSize:  pageSize,
		},
		Search:    queryDTO.Search,
		SortBy:    queryDTO.SortBy,
		SortOrder: queryDTO.SortOrder,
	}

	if queryDTO.Status != "" {
		status := domain.TaskStatus(queryDTO.Status)
		if !domain.IsValidTaskStatus(status) {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid status"))
			return
		}
		query.Status = &status
	}

	tasks, total, err := h.service.GetTasks(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve tasks", err.Error()))
		return
	}

	totalPages := 0
	if pageSize > 0 {
		totalPages = (total + pageSize - 1) / pageSize
	}

	meta := common.PaginationMeta{
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}

	response := dto.TaskListResponse{
		Tasks: tasks,
		Meta:  meta,
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(response))
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid ID"))
		return
	}

	var input dto.TaskUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
		return
	}

	var status *domain.TaskStatus
	if input.Status != nil {
		s := domain.TaskStatus(*input.Status)
		if !domain.IsValidTaskStatus(s) {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid status"))
			return
		}
		status = &s
	}

	task, err := h.service.UpdateTask(c.Request.Context(), uint(id), input.Title, input.Description, status, input.DueAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Failed to update task", err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(task))
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid ID"))
		return
	}

	if err := h.service.DeleteTask(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Failed to delete task", err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewSimpleSuccessResponse("Task deleted"))
}
