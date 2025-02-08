// internal/adapters/db/thumb_test.go
package db

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/mocks/adaptermocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestPostgresThumbRepository_Create(t *testing.T) {
	mockDB := new(adaptermocks.IDB)
	repo := NewPostgresThumbRepository(mockDB)

	t.Run("successful creation", func(t *testing.T) {
		process := &entity.ThumbProcess{
			Video: entity.ThumbProcessVideo{
				Path: "test-video.mp4",
			},
			Thumbnail: entity.ThumbProcessThumb{
				Path: "test-thumb.jpg",
			},
			Status: "pending",
			Error:  "",
		}

		mockDB.On("Create", mock.AnythingOfType("*db.ThumbPostgres")).Return(&gorm.DB{}).Once()

		err := repo.Create(process)

		assert.NoError(t, err)
	})

	t.Run("database error", func(t *testing.T) {
		process := &entity.ThumbProcess{
			Video: entity.ThumbProcessVideo{
				Path: "test-video.mp4",
			},
			Thumbnail: entity.ThumbProcessThumb{
				Path: "test-thumb.jpg",
			},
			Status: "pending",
			Error:  "",
		}

		expectedError := errors.New("database error")
		mockDB.On("Create", mock.AnythingOfType("*db.ThumbPostgres")).Return(&gorm.DB{Error: expectedError}).Once()

		err := repo.Create(process)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestPostgresThumbRepository_List(t *testing.T) {
	mockDB := new(adaptermocks.IDB)
	repo := NewPostgresThumbRepository(mockDB)

	t.Run("successful list retrieval", func(t *testing.T) {
		mockDB.On("Find", mock.AnythingOfType("*[]db.ThumbPostgres"), mock.Anything).
			Run(func(args mock.Arguments) {
				arg := args.Get(0).(*[]ThumbPostgres)
				*arg = []ThumbPostgres{
					{
						BaseModel: BaseModel{
							ID: uuid.New(),
						},
						VideoPath:     "video1.mp4",
						ThumbnailPath: "thumb1.jpg",
						Status:        "completed",
						Error:         "",
					},
					{
						BaseModel: BaseModel{
							ID: uuid.New(),
						},
						VideoPath:     "video2.mp4",
						ThumbnailPath: "thumb2.jpg",
						Status:        "pending",
						Error:         "",
					},
				}
			}).Return(&gorm.DB{}).Once()

		processes := repo.List()

		assert.Len(t, processes, 2)
		assert.Equal(t, "video1.mp4", processes[0].Video.Path)
		assert.Equal(t, "thumb1.jpg", processes[0].Thumbnail.Path)
		assert.Equal(t, "completed", processes[0].Status)
		assert.Equal(t, "video2.mp4", processes[1].Video.Path)
		assert.Equal(t, "thumb2.jpg", processes[1].Thumbnail.Path)
		assert.Equal(t, "pending", processes[1].Status)
	})

	t.Run("empty list", func(t *testing.T) {
		mockDB.On("Find", mock.AnythingOfType("*[]db.ThumbPostgres"), mock.Anything).
			Run(func(args mock.Arguments) {
				arg := args.Get(0).(*[]ThumbPostgres)
				*arg = []ThumbPostgres{}
			}).Return(&gorm.DB{}).Once()

		processes := repo.List()

		assert.Empty(t, processes)
	})
}

func TestPostgresThumbRepository_Update(t *testing.T) {
	mockDB := new(adaptermocks.IDB)
	repo := NewPostgresThumbRepository(mockDB)

	t.Run("successful update", func(t *testing.T) {
		processID := uuid.New()
		process := &entity.ThumbProcess{
			ID: processID,
			Video: entity.ThumbProcessVideo{
				Path: "updated-video.mp4",
			},
			Thumbnail: entity.ThumbProcessThumb{
				Path: "updated-thumb.jpg",
			},
			Status: "completed",
			Error:  "",
		}

		mockDB.On("Save", mock.AnythingOfType("*db.ThumbPostgres")).Return(&gorm.DB{}).Once()

		err := repo.Update(process)

		assert.NoError(t, err)
	})

	t.Run("update with missing ID", func(t *testing.T) {
		process := &entity.ThumbProcess{
			Video: entity.ThumbProcessVideo{
				Path: "updated-video.mp4",
			},
			Thumbnail: entity.ThumbProcessThumb{
				Path: "updated-thumb.jpg",
			},
			Status: "completed",
			Error:  "",
		}

		err := repo.Update(process)

		assert.Error(t, err)
		assert.Equal(t, "process id is required", err.Error())
		mockDB.AssertNotCalled(t, "Save")
	})

	t.Run("database error during update", func(t *testing.T) {
		processID := uuid.New()
		process := &entity.ThumbProcess{
			ID: processID,
			Video: entity.ThumbProcessVideo{
				Path: "updated-video.mp4",
			},
			Thumbnail: entity.ThumbProcessThumb{
				Path: "updated-thumb.jpg",
			},
			Status: "completed",
			Error:  "",
		}

		expectedError := errors.New("database error")
		mockDB.On("Save", mock.AnythingOfType("*db.ThumbPostgres")).Return(&gorm.DB{Error: expectedError}).Once()

		err := repo.Update(process)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestNewPostgresThumbRepository(t *testing.T) {
	mockDB := new(adaptermocks.IDB)
	repo := NewPostgresThumbRepository(mockDB)

	assert.NotNil(t, repo)
	assert.Equal(t, mockDB, repo.db)
}

func TestThumbPostgres_TableName(t *testing.T) {
	thumb := ThumbPostgres{}
	assert.Equal(t, "thumb", thumb.TableName())
}
