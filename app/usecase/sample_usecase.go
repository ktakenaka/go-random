package usecase

import (
	"github.com/ktakenaka/go-random/app/domain/entity"
	"github.com/ktakenaka/go-random/app/domain/repository"
	"github.com/ktakenaka/go-random/app/domain/service"
)

type SampleUsecase struct {
	repo repository.SampleRepository
	srv  *service.SampleService
}

func NewSampleUsecase(repo repository.SampleRepository, srv *service.SampleService) *SampleUsecase {
	return &SampleUsecase{
		repo: repo,
		srv:  srv,
	}
}

func (s *SampleUsecase) ListSample() ([]*entity.Sample, error) {
	samples, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return samples, nil
}

func (s *SampleUsecase) FindSample(id int64) (*entity.Sample, error) {
	sample, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return sample, nil
}

func (s *SampleUsecase) RegisterSample(title string) error {
	if err := s.srv.Duplicated(title); err != nil {
		return err
	}

	if err := s.repo.Create(title); err != nil {
		return err
	}
	return nil
}

func (s *SampleUsecase) UpdateSample(id int64, title string) error {
	if err := s.srv.Duplicated(title); err != nil {
		return err
	}

	if err := s.repo.Update(id, title); err != nil {
		return err
	}
	return nil
}

func (s *SampleUsecase) DeleteSample(id int64) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
