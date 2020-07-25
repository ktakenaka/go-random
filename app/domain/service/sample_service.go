package service

import (
	"fmt"

	"github.com/ktakenaka/go-random/app/domain/repository"
)

type SampleService struct {
	repo repository.SampleRepository
}

func NewSampleService(repo repository.SampleRepository) *SampleService {
	return &SampleService{
		repo: repo,
	}
}

func (s *SampleService) Duplicated(title string) error {
	sample, _ := s.repo.FindByTitle(title)
	if sample != nil {
		return fmt.Errorf("%s already exists", title)
	}

	return nil
}
