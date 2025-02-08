package contracts

import "github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"

type IThumbRepositoryAdapter interface {
	Create(process *entity.ThumbProcess) error
	Update(process *entity.ThumbProcess) (*entity.ThumbProcess, error)
	List() *[]entity.ThumbProcess
}
