package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/middleware"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/ports"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/thumb"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/users"
	"net/http"
)

type ThumbHandler struct {
	thumbService thumb.IThumbService
}

func NewThumbHandler(service thumb.IThumbService) *ThumbHandler {
	return &ThumbHandler{
		thumbService: service,
	}
}

func (h *ThumbHandler) RegisterRoutes(router *gin.RouterGroup) {
	thumbGroup := router.Group("/thumbs")
	thumbGroup.POST("", middleware.AuthMiddleware(users.GetAllowedUsers), h.CreateProcess)
	thumbGroup.PUT("/:id", h.UpdateProcess)
	thumbGroup.GET("", middleware.AuthMiddleware(users.GetAllowedUsers), h.ListProcesses)
}

// @Summary Create a new thumbnail process
// @Description Start a new asynchronous thumbnail generation process from S3 video URL
// @Tags Video Thumbs Processor
// @Accept json
// @Produce json
// @Param request body CreateProcessRequest true "Video URL"
// @Success 202 {string} string "Process started"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /thumbs [post]
func (h *ThumbHandler) CreateProcess(c *gin.Context) {
	var request CreateProcessRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
		})
		return
	}

	if request.URL == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "URL is required",
		})
		return
	}

	var userEmail string
	ctxUser, ok := c.Get("user")
	if ok {
		userEmail = ctxUser.(string)
	}
	err := h.thumbService.CreateProcessAsync(&ports.CreateProcessRequest{
		UserEmail: userEmail,
		Url:       request.URL,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Process started successfully",
	})
}

// @Summary Update a thumbnail process
// @Description Update the status of an existing thumbnail process
// @Tags Video Thumbs Processor
// @Accept json
// @Produce json
// @Param id path string true "Process ID"
// @Param request body UpdateProcessRequest true "Process update information"
// @Success 200 {object} ThumbProcessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /thumbs/{id} [put]
func (h *ThumbHandler) UpdateProcess(c *gin.Context) {
	var request UpdateProcessRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
		})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Process ID is required",
		})
		return
	}

	updated, err := h.thumbService.UpdateProcess(&ports.UpdateProcessRequest{
		ID:            id,
		Status:        request.Status,
		Error:         request.Error,
		ThumbnailPath: request.ThumbnailPath,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ThumbProcessResponse{
		ID:            updated.ID.String(),
		Status:        updated.Status,
		Error:         updated.Error,
		ThumbnailPath: updated.Thumbnail.Path,
	})
}

// @Summary List all thumbnail processes
// @Description Get a list of all thumbnail processes
// @Tags Video Thumbs Processor
// @Produce json
// @Success 200 {array} ThumbProcessResponse
// @Failure 500 {object} ErrorResponse
// @Router /thumbs [get]
func (h *ThumbHandler) ListProcesses(c *gin.Context) {
	processes := h.thumbService.ListProcess()

	response := make([]ThumbProcessResponse, len(*processes))
	for i, process := range *processes {
		response[i] = ThumbProcessResponse{
			ID:            process.ID.String(),
			Status:        process.Status,
			Error:         process.Error,
			ThumbnailPath: process.Thumbnail.Path,
		}
	}

	c.JSON(http.StatusOK, response)
}

type CreateProcessRequest struct {
	URL string `json:"url" binding:"required"`
}

type UpdateProcessRequest struct {
	Status        string `json:"status" binding:"required"`
	Error         string `json:"error,omitempty"`
	ThumbnailPath string `json:"thumbnail_path,omitempty"`
}

type ThumbProcessResponse struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	Error         string `json:"error,omitempty"`
	ThumbnailPath string `json:"thumbnail_path,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
