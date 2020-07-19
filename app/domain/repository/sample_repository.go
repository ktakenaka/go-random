package repository

import (
	"github.com/ktakenaka/go-random/app/domain/entity"
)

type SampleRepository interface {
	FindAll() ([]*entity.Sample, error)
	FindByID(id int) (*entity.Sample, error)
}
