package usecase

import (
	"github.com/jinzhu/copier"

	"github.com/ktakenaka/go-random/app/domain/entity"
	"github.com/ktakenaka/go-random/app/domain/repository"
	"github.com/ktakenaka/go-random/app/domain/service"
	"github.com/ktakenaka/go-random/app/usecase/dto"
)

type SampleUsecase struct {
	repo repository.SampleRepository
	txm  repository.TransactionManager
	srv  *service.SampleService
}

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

func (s *SampleUsecase) ListSample(userID uint64) ([]*entity.Sample, error) {
	// TODO: refactor to use gorm association
	samples, err := s.repo.FindAll(userID)
	if err != nil {
		return nil, err
	}
	return samples, nil
}

func (s *SampleUsecase) FindSample(userID, id uint64) (*entity.Sample, error) {
	sample, err := s.repo.FindByID(userID, id)
	if err != nil {
		return nil, err
	}
	return sample, nil
}

func (s *SampleUsecase) RegisterSample(req dto.CreateSample) error {
	var sample entity.Sample
	if err := copier.Copy(&sample, &req); err != nil {
		return err
	}

	// TODO: enable validation
	// if err := s.srv.Duplicated(sample); err != nil {
	// 	return err
	// }
	err := s.repo.Create(&sample)
	return err
}

func (s *SampleUsecase) UpdateSample(req dto.UpdateSample) (err error) {
	s.beginTx()
	defer func() {
		err = s.endTx(err)
	}()

	var sample entity.Sample
	if err := copier.Copy(&sample, &req); err != nil {
		return err
	}

	// TODO: enable validation
	// if err := s.srv.Duplicated(sample); err != nil {
	// 	return err
	// }

	err = s.repo.Update(&sample)
	return err
}

func (s *SampleUsecase) DeleteSample(userID, id uint64) error {
	err := s.repo.Delete(userID, id)
	return err
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
