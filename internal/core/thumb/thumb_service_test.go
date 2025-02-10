package thumb

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/ports"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/mocks/adaptermocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProcessAsync(t *testing.T) {
	testAppRequest := &ports.CreateProcessRequest{Url: "https://example.com/video.mp4"}

	mockRepo := new(adaptermocks.IThumbRepositoryAdapter)
	mockQueue := new(adaptermocks.IThumbQueueAdapter)
	service := NewThumbService(mockRepo, mockQueue)

	t.Run("successful creation", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*entity.ThumbProcess")).Return(nil).Once()
		mockQueue.On("SendEvent", mock.AnythingOfType("*entity.ThumbProcess")).Return(nil).Once()

		err := service.CreateProcessAsync(testAppRequest)

		assert.NoError(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*entity.ThumbProcess")).Return(errors.New("db error")).Once()

		err := service.CreateProcessAsync(testAppRequest)

		assert.Error(t, err)
	})

	t.Run("queue error", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*entity.ThumbProcess")).Return(nil)
		mockQueue.On("SendEvent", mock.AnythingOfType("*entity.ThumbProcess")).Return(errors.New("queue error")).Once()

		err := service.CreateProcessAsync(testAppRequest)

		assert.Error(t, err)
	})
}

func TestListProcess(t *testing.T) {
	mockRepo := new(adaptermocks.IThumbRepositoryAdapter)
	mockQueue := new(adaptermocks.IThumbQueueAdapter)
	service := NewThumbService(mockRepo, mockQueue)

	expectedList := &[]entity.ThumbProcess{
		*entity.NewThumbProcess("https://example.com/video1.mp4", "teste@teste.com"),
		*entity.NewThumbProcess("https://example.com/video2.mp4", "teste@teste.com"),
	}

	mockRepo.On("List").Return(expectedList)
	result := service.ListProcess()

	assert.Equal(t, expectedList, result)
}

func TestUpdateProcess(t *testing.T) {
	mockRepo := new(adaptermocks.IThumbRepositoryAdapter)
	mockQueue := new(adaptermocks.IThumbQueueAdapter)
	service := NewThumbService(mockRepo, mockQueue)
	mockedID, _ := uuid.NewV7()

	request := &ports.UpdateProcessRequest{
		ID:            mockedID,
		Status:        entity.ThumbProcessStatusComplete,
		Error:         "",
		ThumbnailPath: "https://example.com/video-thumbs.zip",
	}

	thumbProcess := &entity.ThumbProcess{
		ID:        request.ID,
		Status:    request.Status,
		Error:     request.Error,
		Thumbnail: entity.ThumbProcessThumb{Path: request.ThumbnailPath},
	}

	t.Run("successful creation", func(t *testing.T) {
		mockRepo.On("Update", thumbProcess).Return(thumbProcess, nil).Once()

		updated, err := service.UpdateProcess(request)

		assert.NotNil(t, updated)
		assert.NoError(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo.On("Update", mock.AnythingOfType("*entity.ThumbProcess")).Return(nil, errors.New("db error")).Once()

		updated, err := service.UpdateProcess(request)
		assert.Nil(t, updated)
		assert.Error(t, err)
	})
}
