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

func (s *SampleService) Duplicated(userID uint64, title string) error {
	sample, _ := s.repo.FindByTitle(userID, title)
	if sample != nil {
		return fmt.Errorf("%s already exists", title)
	}

	return nil
}
