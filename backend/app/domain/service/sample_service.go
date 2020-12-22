package service

import (
	"fmt"

	"github.com/ktakenaka/go-random/backend/app/domain/repository"
)

// SampleService service
type SampleService struct {
	repo repository.SampleRepository
}

// NewSampleService constructor
func NewSampleService(repo repository.SampleRepository) *SampleService {
	return &SampleService{
		repo: repo,
	}
}

// Duplicated check if it's not duplicated
func (s *SampleService) Duplicated(userID, title string) error {
	_, err := s.repo.FindByTitle(userID, title)
	if err != nil {
		return fmt.Errorf("%s already exists", title)
	}

	return nil
}
