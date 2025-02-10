package contracts

import (
	"context"

	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"
)

type IThumbRepositoryAdapter interface {
	Create(ctx context.Context, process *entity.ThumbProcess) error
	Update(ctx context.Context, process *entity.ThumbProcess) (*entity.ThumbProcess, error)
	List(ctx context.Context) *[]entity.ThumbProcess
}
