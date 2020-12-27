package repository

import (
	"github.com/ktakenaka/go-random/backend/app/domain/entity"
)

// SampleRepository interface for sample
type SampleRepository interface {
	FindAll(uesrID string, query *entity.SampleQuery) ([]entity.Sample, error)
	FindByID(uesrID, id string) (entity.Sample, error)
	FindByTitle(userID, title string) (entity.Sample, error)
	Create(sample *entity.Sample) (*entity.Sample, error)
	Update(sample *entity.Sample) (*entity.Sample, error)
	Delete(userID, id string) error
}
