package repository

import (
	"github.com/ktakenaka/go-random/backend/app/domain/entity"
)

// SampleRepository interface for sample
type SampleRepository interface {
	FindAll(uesrID uint64) ([]entity.Sample, error)
	FindByID(uesrID, id uint64) (entity.Sample, error)
	FindByTitle(userID uint64, title string) (entity.Sample, error)
	Create(sample *entity.Sample) (*entity.Sample, error)
	Update(sample *entity.Sample) (*entity.Sample, error)
	Delete(userID, id uint64) error
	AssignTx(txm TransactionManager)
}
