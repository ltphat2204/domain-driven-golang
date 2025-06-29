package handler

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/domain-driven-golang/modules/category/application"
	"github.com/ltphat2204/domain-driven-golang/modules/category/domain"
	"github.com/ltphat2204/domain-driven-golang/modules/category/dto"
	"github.com/ltphat2204/domain-driven-golang/common"
)

type CategoryHandler struct {
	application application.CategoryService
}

func NewCategoryHandler(application application.CategoryService) *CategoryHandler {
	return &CategoryHandler{application: application}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var input dto.CategoryCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
		return
	}

	category, err := h.application.CreateCategory(c.Request.Context(), input.Name, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Failed to create category", err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(category))
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid ID"))
		return
	}

	category, err := h.application.GetCategoryByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewErrorResponse(http.StatusNotFound, "Category not found", err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(category))
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	var queryDTO dto.CategoryQueryDTO
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

	allowedSortFields := []string{"name", "created_at"}
	allowedSortOrders := []string{"asc", "desc"}

	if queryDTO.SortBy != "" && !slices.Contains(allowedSortFields, queryDTO.SortBy) {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid sort_by"))
		return
	}

	if queryDTO.SortOrder != "" && !slices.Contains(allowedSortOrders, queryDTO.SortOrder) {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid sort_order"))
		return
	}

	query := &domain.CategoryQuery{
		BaseQuery: common.BaseQuery{
			Page:     page,
			PageSize: pageSize,
		},
		Search:    queryDTO.Search,
		SortBy:    queryDTO.SortBy,
		SortOrder: queryDTO.SortOrder,
	}

	categories, total, err := h.application.GetCategories(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve categories", err.Error()))
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

	response := dto.CategoryListResponse{
		Categories: categories,
		Meta:       meta,
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(response))
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid ID"))
		return
	}

	var input dto.CategoryUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
		return
	}

	category, err := h.application.UpdateCategory(c.Request.Context(), uint(id), input.Name, input.Description, input.Color)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(http.StatusBadRequest, "Failed to update category", err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(category))
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("Invalid ID"))
		return
	}

	if err := h.application.DeleteCategory(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Failed to delete category", err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewSimpleSuccessResponse("Category deleted"))
}
