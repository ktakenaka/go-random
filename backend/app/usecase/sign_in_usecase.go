package usecase

import (
	"context"
	"time"

	"github.com/ktakenaka/go-random/backend/app/domain/repository"
	appErr "github.com/ktakenaka/go-random/backend/app/errors"
	"github.com/ktakenaka/go-random/backend/pkg/jwtutil"
	"golang.org/x/xerrors"
)

// SignInUsecase usecase
type SignInUsecase struct {
	googleRepo repository.GoogleRepository
	userRepo   repository.UserRepository
}

// NewSignInUsecase constructor
func NewSignInUsecase(
	gRepo repository.GoogleRepository, uRepo repository.UserRepository,
) *SignInUsecase {
	return &SignInUsecase{
		googleRepo: gRepo,
		userRepo:   uRepo,
	}
}

// Execute sign in flow
func (uc *SignInUsecase) Execute(code, nonce string) (aTkn, rTkn, csrfTkn string, err error) {
	ctx := context.Background()

	token, err := uc.googleRepo.GetToken(ctx, code, nonce)
	if err != nil {
		err = appErr.NewAppError(err)
		return "", "", "", xerrors.Errorf("%w", err)
	}

	body, err := uc.googleRepo.GetUserInfo(ctx, token)
	if err != nil {
		err = appErr.NewAppError(err)
		return "", "", "", xerrors.Errorf("%w", err)
	}

	// TODO: put User entity
	user, err := uc.userRepo.UpdateOrCreate(body)
	if err != nil {
		err = appErr.NewAppError(err)
		return "", "", "", xerrors.Errorf("%w", err)
	}

	csrfTkn, err = jwtutil.GenerateCSRFToken()
	if err != nil {
		err = appErr.NewAppError(err)
		return "", "", "", xerrors.Errorf("%w", err)
	}

	claims := jwtutil.AuthClaims{
		UserID:    user.ID.String(),
		IssueTime: time.Now(),
		CSRFToken: csrfTkn,
	}

	aTkn, rTkn, err = jwtutil.GenerateToken(&claims)
	if err != nil {
		err = appErr.NewAppError(err)
		return "", "", "", xerrors.Errorf("%w", err)
	}
	return aTkn, rTkn, csrfTkn, err
}
