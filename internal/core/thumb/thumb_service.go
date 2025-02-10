package thumb

import (
	"context"
	"errors"

	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/contracts"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/ports"
)

type IThumbService interface {
	CreateProcessAsync(ctx context.Context, request *ports.CreateProcessRequest) error
	UpdateProcess(ctx context.Context, request *ports.UpdateProcessRequest) (*entity.ThumbProcess, error)
	ListProcess(ctx context.Context) *[]entity.ThumbProcess
}

type Service struct {
	processRepository contracts.IThumbRepositoryAdapter
	queueAdapter      contracts.IThumbQueueAdapter
}

func NewThumbService(repo contracts.IThumbRepositoryAdapter, q contracts.IThumbQueueAdapter) *Service {
	return &Service{processRepository: repo, queueAdapter: q}
}

func (s *Service) CreateProcessAsync(ctx context.Context, request *ports.CreateProcessRequest) error {
	thumbProcess := entity.NewThumbProcess(request.Url, request.UserEmail)

	if err := s.processRepository.Create(ctx, thumbProcess); err != nil {
		return err
	}

	if err := s.queueAdapter.SendEvent(ctx, thumbProcess); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateProcess(ctx context.Context, request *ports.UpdateProcessRequest) (*entity.ThumbProcess, error) {
	if !entity.AllowedProcessStatus[request.Status] {
		return nil, errors.New("informed status not allowed")
	}

	thumbProcess := &entity.ThumbProcess{
		ID:        request.ID,
		Status:    request.Status,
		Error:     request.Error,
		Thumbnail: entity.ThumbProcessThumb{Path: request.ThumbnailPath},
	}

	updated, err := s.processRepository.Update(ctx, thumbProcess)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *Service) ListProcess(ctx context.Context) *[]entity.ThumbProcess {
	processList := s.processRepository.List(ctx)

	return processList
}
