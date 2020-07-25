package repository

import (
	"github.com/ktakenaka/go-random/app/domain/entity"
)

type SampleRepository interface {
	FindAll() ([]*entity.Sample, error)
	FindByID(id int64) (*entity.Sample, error)
	FindByTitle(title string) (*entity.Sample, error)
	Create(title string) error
	Update(id int64, title string) error
	Delete(id int64) error
}
