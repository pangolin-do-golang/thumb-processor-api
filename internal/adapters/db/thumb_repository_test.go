// internal/adapters/db/thumb_test.go
package db

import (
	"context"
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
	mockDB, repo, _, ctx := setupTest()

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

		err := repo.Create(ctx, process)

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

		err := repo.Create(ctx, process)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("create with missing user ID", func(t *testing.T) {
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

		err := repo.Create(context.Background(), process)

		assert.Error(t, err)
		assert.Equal(t, RequiredUserIDError, err)
	})
}

func TestPostgresThumbRepository_List(t *testing.T) {
	mockDB, repo, mockedUserID, ctx := setupTest()

	t.Run("successful list retrieval", func(t *testing.T) {
		mockDB.On("Find", mock.AnythingOfType("*[]db.ThumbPostgres"), "user_id = ?", mockedUserID).
			Run(func(args mock.Arguments) {
				arg := args.Get(0).(*[]ThumbPostgres)
				*arg = []ThumbPostgres{
					{
						BaseModel: BaseModel{
							ID: uuid.New(),
						},
						UserID:        1,
						VideoPath:     "video1.mp4",
						ThumbnailPath: "thumb1.jpg",
						Status:        "completed",
						Error:         "",
					},
					{
						BaseModel: BaseModel{
							ID: uuid.New(),
						},
						UserID:        1,
						VideoPath:     "video2.mp4",
						ThumbnailPath: "thumb2.jpg",
						Status:        "pending",
						Error:         "",
					},
				}
			}).Return(&gorm.DB{}).Once()

		processes := *repo.List(ctx)

		assert.Len(t, processes, 2)
		assert.Equal(t, "video1.mp4", processes[0].Video.Path)
		assert.Equal(t, "thumb1.jpg", processes[0].Thumbnail.Path)
		assert.Equal(t, "completed", processes[0].Status)
		assert.Equal(t, "video2.mp4", processes[1].Video.Path)
		assert.Equal(t, "thumb2.jpg", processes[1].Thumbnail.Path)
		assert.Equal(t, "pending", processes[1].Status)
	})

	t.Run("empty list", func(t *testing.T) {
		mockDB.On("Find", mock.AnythingOfType("*[]db.ThumbPostgres"), "user_id = ?", mockedUserID).
			Run(func(args mock.Arguments) {
				arg := args.Get(0).(*[]ThumbPostgres)
				*arg = []ThumbPostgres{}
			}).Return(&gorm.DB{}).Once()

		processes := repo.List(ctx)

		assert.Empty(t, processes)
	})

	t.Run("list with missing user ID", func(t *testing.T) {
		processes := repo.List(context.Background())
		assert.Empty(t, processes)
	})
}

func TestPostgresThumbRepository_Update(t *testing.T) {
	mockDB, repo, _, ctx := setupTest()

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

		mockDB.On("Updates", mock.AnythingOfType("*db.ThumbPostgres")).Return(&gorm.DB{}).Once()

		updated, err := repo.Update(ctx, process)

		assert.NoError(t, err)
		assert.NotNil(t, updated)
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

		updated, err := repo.Update(ctx, process)

		assert.Nil(t, updated)
		assert.Error(t, err)
		assert.Equal(t, "process id is required", err.Error())
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
		mockDB.On("Updates", mock.AnythingOfType("*db.ThumbPostgres")).Return(&gorm.DB{Error: expectedError}).Once()

		updated, err := repo.Update(ctx, process)

		assert.Nil(t, updated)
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

func setupTest() (*adaptermocks.IDB, *PostgresThumbRepository, int, context.Context) {
	mockedUserID := 1
	mockDB := new(adaptermocks.IDB)

	return mockDB,
		NewPostgresThumbRepository(mockDB),
		mockedUserID,
		context.WithValue(context.Background(), "logged_user_id", mockedUserID)
}
