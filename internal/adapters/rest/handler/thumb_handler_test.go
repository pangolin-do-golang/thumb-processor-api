package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/mocks/servicemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTest() (*gin.Engine, *servicemocks.IThumbService) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockService := new(servicemocks.IThumbService)
	handler := NewThumbHandler(mockService)
	handler.RegisterRoutes(router)
	return router, mockService
}

func TestCreateProcess(t *testing.T) {
	router, mockService := setupTest()

	t.Run("successful creation", func(t *testing.T) {
		mockService.On("CreateProcessAsync", mock.AnythingOfType("*ports.CreateProcessRequest")).Return(nil).Once()

		body := CreateProcessRequest{URL: "https://example.com/video.mp4"}
		jsonBody, _ := json.Marshal(body)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/thumbs", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusAccepted, w.Code)
	})

	t.Run("invalid request", func(t *testing.T) {
		body := CreateProcessRequest{URL: ""} // Empty URL
		jsonBody, _ := json.Marshal(body)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/thumbs", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUpdateProcess(t *testing.T) {
	router, mockService := setupTest()
	mockedUUID, _ := uuid.NewV7()

	t.Run("successful update", func(t *testing.T) {
		updatedProcess := &entity.ThumbProcess{
			ID:     mockedUUID,
			Status: "completed",
			Thumbnail: entity.ThumbProcessThumb{
				Path: "path/to/thumbnail.jpg",
			},
		}

		mockService.On("UpdateProcess", mock.AnythingOfType("*ports.UpdateProcessRequest")).Return(updatedProcess, nil).Once()

		body := UpdateProcessRequest{
			Status:        "completed",
			ThumbnailPath: "path/to/thumbnail.jpg",
		}
		jsonBody, _ := json.Marshal(body)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/thumbs/"+mockedUUID.String(), bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestListProcesses(t *testing.T) {
	router, mockService := setupTest()
	mockedUUIDComplete, _ := uuid.NewV7()
	mockedUUIDProcessing, _ := uuid.NewV7()

	processes := &[]entity.ThumbProcess{
		{
			ID:     mockedUUIDComplete,
			Video:  entity.ThumbProcessVideo{Path: "https://example.com/video1.mp4"},
			Status: "completed",
		},
		{
			ID:     mockedUUIDProcessing,
			Video:  entity.ThumbProcessVideo{Path: "https://example.com/video2.mp4"},
			Status: "processing",
		},
	}

	mockService.On("ListProcess").Return(processes).Once()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/thumbs", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []ThumbProcessResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
}
