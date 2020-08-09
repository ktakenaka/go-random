package usecase

import (
	"github.com/ktakenaka/go-random/app/domain/repository"
)

type SignInUsecase struct {
	googleRepo repository.GoogleRepository
	userRepo   repository.UserRepository
}

func NewSignInUsecase(
	gRepo repository.GoogleRepository, uRepo repository.UserRepository,
) *SignInUsecase {
	return &SignInUsecase{
		googleRepo: gRepo,
		userRepo:   uRepo,
	}
}

func (uc *SignInUsecase) Execute(code string) (interface{}, error) {
	token, err := uc.googleRepo.GetToken(code)
	if err != nil {
		return nil, err
	}

	body, err := uc.googleRepo.GetUserInfo(token)
	if err != nil {
		return nil, err
	}

	return body, nil
}
