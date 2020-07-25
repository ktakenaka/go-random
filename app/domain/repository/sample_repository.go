package repository

import (
	"github.com/ktakenaka/go-random/app/domain/entity"
)

type SampleRepository interface {
	FindAll() ([]*entity.Sample, error)
	FindByID(id string) (*entity.Sample, error)
	FindByTitle(title string) (*entity.Sample, error)
	Create(title string) error
	Update(id, title string) error
	Delete(id string) error
}
