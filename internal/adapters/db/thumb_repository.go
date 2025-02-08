package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"
)

type ThumbPostgres struct {
	BaseModel
	VideoPath     string
	Status        string
	Error         string
	ThumbnailPath string
}

func (op ThumbPostgres) TableName() string {
	return "thumb"
}

type PostgresThumbRepository struct {
	db IDB
}

func (r *PostgresThumbRepository) Create(process *entity.ThumbProcess) error {
	record := &ThumbPostgres{
		VideoPath:     process.Video.Path,
		ThumbnailPath: process.Thumbnail.Path,
		Status:        process.Status,
		Error:         process.Error,
	}

	result := r.db.Create(record)

	return result.Error
}

func (r *PostgresThumbRepository) List() *[]entity.ThumbProcess {
	processes := []entity.ThumbProcess{}
	records := []ThumbPostgres{}

	r.db.Find(&records)

	for _, record := range records {
		processes = append(processes, entity.ThumbProcess{
			ID:     record.ID,
			Status: record.Status,
			Error:  record.Error,
			Video: entity.ThumbProcessVideo{
				Path: record.VideoPath,
			},
			Thumbnail: entity.ThumbProcessThumb{
				Path: record.ThumbnailPath,
			},
		})
	}

	return &processes
}

func NewPostgresThumbRepository(db IDB) *PostgresThumbRepository {
	return &PostgresThumbRepository{db: db}
}

func (r *PostgresThumbRepository) Update(process *entity.ThumbProcess) (*entity.ThumbProcess, error) {
	if process.ID == uuid.Nil {
		return nil, errors.New("process id is required")
	}

	result := r.db.Save(&ThumbPostgres{
		BaseModel: BaseModel{
			ID: process.ID,
		},
		VideoPath:     process.Video.Path,
		ThumbnailPath: process.Thumbnail.Path,
		Status:        process.Status,
		Error:         process.Error,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity.ThumbProcess{
		ID:     process.ID,
		Status: process.Status,
		Error:  process.Error,
		Video: entity.ThumbProcessVideo{
			Path: process.Video.Path,
		},
		Thumbnail: entity.ThumbProcessThumb{
			Path: process.Thumbnail.Path,
		},
	}, nil
}
