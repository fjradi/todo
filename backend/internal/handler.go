package internal

import (
	"github.com/gin-gonic/gin"
	"my-app/internal/repo"
	"net/http"
)

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

type AddTodoRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) AddTodo(ctx *gin.Context) {
	var request AddTodoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": parseValidationError(err)})
		return
	}

	todo, err := h.service.AddTodo(ctx.Request.Context(), request.Name)
	if err != nil {
		statusCode, message := parseSqlError(err)
		ctx.JSON(statusCode, gin.H{"error": message})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (h *Handler) GetTodos(ctx *gin.Context) {
	todos, err := h.service.GetTodos(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

type UpdateTodoUriRequest struct {
	Id string `uri:"id" binding:"required"`
}

type UpdateTodoBodyRequest struct {
	Name        string `json:"name" binding:"required"`
	IsCompleted *bool  `json:"is_completed" binding:"required"`
}

func (h *Handler) UpdateTodo(ctx *gin.Context) {
	var requestUri UpdateTodoUriRequest
	var requestBody UpdateTodoBodyRequest
	if err := ctx.ShouldBindUri(&requestUri); err != nil {
		ctx.JSON(http.StatusBadRequest, parseValidationError(err))
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, parseValidationError(err))
		return
	}

	id, err := parseStringToUUID(requestUri.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": []ApiError{{
			Param:   "id",
			Message: "invalid id format",
		}}})
		return
	}

	todo, err := h.service.UpdateTodo(ctx.Request.Context(), repo.UpdateTodoParams{
		ID:          id,
		Name:        requestBody.Name,
		IsCompleted: *requestBody.IsCompleted,
	})
	if err != nil {
		statusCode, message := parseSqlError(err)
		ctx.JSON(statusCode, gin.H{"error": message})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}
