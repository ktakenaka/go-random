package mysql

import (
	"github.com/ktakenaka/go-random/backend/app/domain/entity"
	"github.com/ktakenaka/go-random/backend/app/domain/repository"
	"github.com/ktakenaka/go-random/backend/app/errors"
	"gorm.io/gorm"
)

// SampleRepository access sample
type SampleRepository struct {
	DB *gorm.DB
}

// NewSampleRepository constructor
func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{DB: db}
}

// FindAll find all samples
func (r *SampleRepository) FindAll(userID string, query *entity.SampleQuery) ([]entity.Sample, error) {
	samples := make([]entity.Sample, 0)

	tx := r.DB.Where(&entity.Sample{UserID: userID})

	columns := []string{"title", "content", "created_at", "updated_at"}
	query.AddWhereClause(columns, tx)

	if order := query.ToOrderBy(columns); order != "" {
		tx.Order(order)
	}

	err := tx.Limit(query.GetLimit()).Offset(query.GetOffset()).Find(&samples).Error

	return samples, err
}

// FindByTitle find one sample from title
func (r *SampleRepository) FindByTitle(userID string, title string) (entity.Sample, error) {
	var sample entity.Sample
	err := r.DB.Where(&entity.Sample{UserID: userID, Title: title}).First(&sample).Error
	return sample, err
}

// FindByID find on sample from id
func (r *SampleRepository) FindByID(userID, id string) (entity.Sample, error) {
	sample := entity.Sample{UserID: userID}
	err := r.DB.Where("user_id=?", sample.UserID).Take(&sample, id).Error
	if err != nil {
		return sample, errors.Wrap("mysql", err)
	}

	return sample, nil
}

// Create creates sample
func (r *SampleRepository) Create(sample *entity.Sample) (*entity.Sample, error) {
	if err := sample.Validate(); err != nil {
		return sample, err
	}

	err := r.DB.Create(sample).Error
	return sample, err
}

// Delete delete
func (r *SampleRepository) Delete(userID, id string) error {
	sample := entity.Sample{UserID: userID}
	err := r.DB.Where("user_id=?", sample.UserID).Delete(sample, id).Error
	return err
}

// Update update
func (r *SampleRepository) Update(sample *entity.Sample) (*entity.Sample, error) {
	if err := sample.Validate(); err != nil {
		return sample, err
	}

	err := r.DB.Model(sample).Where("user_id = ?", sample.UserID).Updates(&sample).Error
	return sample, err
}

// AssignTx assign transaction
func (r *SampleRepository) AssignTx(txm repository.TransactionManager) {
	r.DB = txm.GetTx()
}
