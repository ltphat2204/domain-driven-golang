package handlers

import (
	"net/http"
	"strconv"
	"github.com/ltphat2204/domain-driven-golang/application"
	"github.com/ltphat2204/domain-driven-golang/common"
	"github.com/ltphat2204/domain-driven-golang/domain"
	"github.com/ltphat2204/domain-driven-golang/dto"
	"github.com/gin-gonic/gin"
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

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve tasks", err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(tasks))
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