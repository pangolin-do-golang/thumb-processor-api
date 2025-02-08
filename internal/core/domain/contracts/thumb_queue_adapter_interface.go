package contracts

import "github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"

type IThumbQueueAdapter interface {
	SendEvent(process *entity.ThumbProcess) error
}
