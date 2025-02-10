package thumb

import (
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/contracts"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/ports"
)

type IThumbService interface {
	CreateProcessAsync(request *ports.CreateProcessRequest) error
	UpdateProcess(request *ports.UpdateProcessRequest) (*entity.ThumbProcess, error)
	ListProcess() *[]entity.ThumbProcess
}

type Service struct {
	processRepository contracts.IThumbRepositoryAdapter
	queueAdapter      contracts.IThumbQueueAdapter
}

func NewThumbService(repo contracts.IThumbRepositoryAdapter, q contracts.IThumbQueueAdapter) *Service {
	return &Service{processRepository: repo, queueAdapter: q}
}

func (s *Service) CreateProcessAsync(request *ports.CreateProcessRequest) error {
	thumbProcess := entity.NewThumbProcess(request.Url, request.UserEmail)

	if err := s.processRepository.Create(thumbProcess); err != nil {
		return err
	}

	if err := s.queueAdapter.SendEvent(thumbProcess); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateProcess(request *ports.UpdateProcessRequest) (*entity.ThumbProcess, error) {
	thumbProcess := &entity.ThumbProcess{
		ID:        request.ID,
		Status:    request.Status,
		Error:     request.Error,
		Thumbnail: entity.ThumbProcessThumb{Path: request.ThumbnailPath},
	}

	updated, err := s.processRepository.Update(thumbProcess)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *Service) ListProcess() *[]entity.ThumbProcess {
	processList := s.processRepository.List()

	return processList
}
