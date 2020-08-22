package usecase

import (
	"time"

	"github.com/ktakenaka/go-random/app/domain/repository"
	"github.com/ktakenaka/go-random/helper/jwtutil"
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

func (uc *SignInUsecase) Execute(code string) (aTkn, rTkn, csrfTkn string, err error) {
	token, err := uc.googleRepo.GetToken(code)
	if err != nil {
		return "", "", "", err
	}

	body, err := uc.googleRepo.GetUserInfo(token)
	if err != nil {
		return "", "", "", err
	}

	user, err := uc.userRepo.UpdateOrCreate(body)
	if err != nil {
		return "", "", "", err
	}

	csrfTkn, err = jwtutil.GenerateCSRFToken()
	if err != nil {
		return "", "", "", err
	}

	claims := jwtutil.AuthClaims{
		UserID:    user.ID, //TODO: make user_id hash
		IssueTime: time.Now(),
		CSRFToken: csrfTkn,
	}

	aTkn, rTkn, csrfTkn, err = jwtutil.GenerateToken(&claims)
	return aTkn, rTkn, csrfTkn, err
}
