package mysql

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/ktakenaka/go-random/app/domain/entity"
)

type SampleRepository struct {
	DB *gorm.DB
}

func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{DB: db}
}

func (r *SampleRepository) FindAll() ([]*entity.Sample, error) {
	samples := make([]*entity.Sample, 0)
	err := r.DB.Find(&samples).Error
	return samples, err
}

func (r *SampleRepository) FindByTitle(title string) (*entity.Sample, error) {
	var sample entity.Sample
	err := r.DB.Where("title = ?", title).First(&sample).Error

	if err != nil {
		return nil, err
	}

	return &sample, nil
}

func (r *SampleRepository) FindByID(id int64) (*entity.Sample, error) {
	var sample entity.Sample
	err := r.DB.First(&sample, id).Error

	if err != nil {
		return nil, err
	}

	return &sample, nil
}

func (r *SampleRepository) Create(title string) error {
	sample := entity.Sample{Title: title}

	r.DB.Create(&sample)
	if r.DB.NewRecord(sample) {
		return fmt.Errorf("failed")
	}
	return nil
}

func (r *SampleRepository) Delete(id int64) error {
	var sample entity.Sample
	if err := r.DB.First(&sample, id).Error; err != nil {
		return err
	}

	if err := r.DB.Delete(&sample).Error; err != nil {
		return err
	}
	return nil
}

func (r *SampleRepository) Update(id int64, title string) error {
	var sample entity.Sample
	if err := r.DB.First(&sample, id).Error; err != nil {
		return err
	}

	sample.Title = title
	sample.UpdatedAt = time.Now()

	if err := r.DB.Save(&sample).Error; err != nil {
		return err
	}
	return nil
}
