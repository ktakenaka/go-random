package mysql

import (
	"github.com/ktakenaka/go-random/backend/app/domain/entity"
	"github.com/ktakenaka/go-random/backend/app/domain/repository"
	"gorm.io/gorm"
)

type SampleRepository struct {
	DB *gorm.DB
}

func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{DB: db}
}

func (r *SampleRepository) FindAll(userID uint64) ([]*entity.Sample, error) {
	samples := make([]*entity.Sample, 0)
	err := r.DB.Where(&entity.Sample{UserID: userID}).Find(&samples).Error
	return samples, err
}

func (r *SampleRepository) FindByTitle(userID uint64, title string) (*entity.Sample, error) {
	var sample entity.Sample
	err := r.DB.Where(&entity.Sample{UserID: userID, Title: title}).First(&sample).Error
	return &sample, err
}

func (r *SampleRepository) FindByID(userID, id uint64) (*entity.Sample, error) {
	var sample entity.Sample
	err := r.DB.Where(&entity.Sample{ID: id, UserID: userID}).First(&sample).Error

	return &sample, err
}

func (r *SampleRepository) Create(sample *entity.Sample) error {
	if err := sample.Validate(); err != nil {
		return err
	}

	err := r.DB.Create(sample).Error
	return err
}

func (r *SampleRepository) Delete(userID, id uint64) error {
	err := r.DB.Where("user_id = ?", userID).Delete(&entity.Sample{ID: id}).Error
	return err
}

func (r *SampleRepository) Update(sample *entity.Sample) error {
	if err := sample.Validate(); err != nil {
		return err
	}

	err := r.DB.Model(sample).Where("user_id = ?", sample.UserID).Updates(&sample).Error
	return err
}

func (r *SampleRepository) AssignTx(txm repository.TransactionManager) {
	r.DB = txm.GetTx()
}
