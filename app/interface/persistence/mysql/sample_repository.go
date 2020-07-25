package mysql

import (
	"fmt"
	"time"

	"github.com/ktakenaka/go-random/app/domain/entity"
)

type SampleRepository struct{}

func NewSampleRepository() *SampleRepository {
	return &SampleRepository{}
}

func (r *SampleRepository) FindAll() ([]*entity.Sample, error) {
	db := DBConnection()
	defer db.Close()

	samples := make([]*entity.Sample, 0)
	err := db.Find(&samples).Error
	return samples, err
}

func (r *SampleRepository) FindByTitle(title string) (*entity.Sample, error) {
	db := DBConnection()
	defer db.Close()

	var sample entity.Sample
	err := db.Where("title = ?", title).First(&sample).Error

	if err != nil {
		return nil, err
	}

	return &sample, nil
}

func (r *SampleRepository) FindByID(id string) (*entity.Sample, error) {
	db := DBConnection()
	defer db.Close()

	var sample entity.Sample
	err := db.First(&sample, id).Error

	if err != nil {
		return nil, err
	}

	return &sample, nil
}

func (r *SampleRepository) Create(title string) error {
	db := DBConnection()
	defer db.Close()

	sample := entity.Sample{Title: title}

	db.Create(&sample)
	if db.NewRecord(sample) {
		return fmt.Errorf("failed")
	}
	return nil
}

func (r *SampleRepository) Delete(id string) error {
	db := DBConnection()
	defer db.Close()

	var sample entity.Sample
	if err := db.First(&sample, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&sample).Error; err != nil {
		return err
	}
	return nil
}

func (r *SampleRepository) Update(id, title string) error {
	db := DBConnection()
	defer db.Close()

	var sample entity.Sample
	if err := db.First(&sample, id).Error; err != nil {
		return err
	}

	sample.Title = title
	sample.UpdatedAt = time.Now()

	if err := db.Save(&sample).Error; err != nil {
		return err
	}
	return nil
}
