package usecase

import (
	"context"
	"time"

	"github.com/ktakenaka/go-random/backend/app/domain/repository"
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

func (uc *SignInUsecase) Execute(code, nonce string) (aTkn, rTkn, csrfTkn string, err error) {
	ctx := context.Background()

	token, err := uc.googleRepo.GetToken(ctx, code, nonce)
	if err != nil {
		return "", "", "", err
	}

	body, err := uc.googleRepo.GetUserInfo(ctx, token)
	if err != nil {
		return "", "", "", err
	}

	// TODO: put User entity
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

	aTkn, rTkn, err = jwtutil.GenerateToken(&claims)
	return aTkn, rTkn, csrfTkn, err
}
