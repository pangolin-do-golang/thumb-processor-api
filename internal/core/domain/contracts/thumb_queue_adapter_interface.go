package contracts

import (
	"context"

	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"
)

type IThumbQueueAdapter interface {
	SendEvent(ctx context.Context, process *entity.ThumbProcess) error
}
