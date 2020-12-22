package usecase

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/ktakenaka/go-random/backend/app/domain/entity"
	"github.com/ktakenaka/go-random/backend/app/domain/repository"
	"github.com/ktakenaka/go-random/backend/app/domain/service"
	"github.com/ktakenaka/go-random/backend/app/errors"
	"github.com/ktakenaka/go-random/backend/app/usecase/dto"
)

// SampleUsecase usecase for sample
type SampleUsecase struct {
	repo repository.SampleRepository
	txm  repository.TransactionManager
	srv  *service.SampleService
}

// NewSampleUsecase constructor
func NewSampleUsecase(
	repo repository.SampleRepository,
	txm repository.TransactionManager,
	srv *service.SampleService,
) *SampleUsecase {
	return &SampleUsecase{
		repo: repo,
		txm:  txm,
		srv:  srv,
	}
}

// List list sample
func (s *SampleUsecase) List(userID string, query *dto.JSONAPIQuery) ([]entity.Sample, error) {
	var enQuery entity.SampleQuery
	err := query.Bind(&enQuery)
	if err != nil {
		return nil, errors.Wrap(fmt.Sprintf("query: %v", query), err)
	}

	samples, err := s.repo.FindAll(userID, &enQuery)
	if err != nil {
		return nil, errors.Wrap(fmt.Sprintf("query: %v", enQuery), err)
	}

	return samples, nil
}

// Find find
func (s *SampleUsecase) Find(userID, id string) (entity.Sample, error) {
	sample, err := s.repo.FindByID(userID, id)
	if err != nil {
		return sample, errors.Wrap(fmt.Sprintf("userID: %v, ID: %v", userID, id), err)
	}
	return sample, nil
}

// Create create
func (s *SampleUsecase) Create(req dto.CreateSample) error {
	sample := &entity.Sample{}
	if err := copier.Copy(sample, &req); err != nil {
		return err
	}

	// TODO: enable validation
	// if err := s.srv.Duplicated(sample); err != nil {
	// 	return err
	// }
	_, err := s.repo.Create(sample)
	return err
}

// Update update
func (s *SampleUsecase) Update(req dto.UpdateSample) (err error) {
	sample := &entity.Sample{}
	if err := copier.Copy(sample, &req); err != nil {
		return err
	}

	// TODO: enable validation
	// if err := s.srv.Duplicated(sample); err != nil {
	// 	return err
	// }
	s.beginTx()
	defer func() {
		err = s.endTx(err)
	}()
	_, err = s.repo.Update(sample)
	return err
}

// Delete delete
func (s *SampleUsecase) Delete(userID, id string) error {
	err := s.repo.Delete(userID, id)
	return err
}

// Import csv
func (s *SampleUsecase) Import(samples []dto.ImportSample) error {
	// TODO: Implement insert on duplicated update
	for _, item := range samples {
		fmt.Println(item)
	}

	return nil
}

// ListForExport for csv export
func (s *SampleUsecase) ListForExport(userID string) ([]dto.ExportSample, error) {
	// TODO: refactor to use gorm association
	samples, err := s.repo.FindAll(userID, &entity.SampleQuery{})
	if err != nil {
		return nil, err
	}

	var dtoSamples []dto.ExportSample
	if err := copier.Copy(&dtoSamples, &samples); err != nil {
		return dtoSamples, err
	}

	return dtoSamples, nil
}

func (s *SampleUsecase) beginTx() {
	s.txm.Begin()
	s.repo.AssignTx(s.txm)
}

func (s *SampleUsecase) endTx(err error) error {
	if p := recover(); p != nil {
		s.txm.Rollback()
		panic(p)
	} else if err != nil {
		s.txm.Rollback()
	} else {
		err = s.txm.Commit()
	}
	return err
}
