package repository

import (
	"golang.org/x/oauth2"

	"github.com/ktakenaka/go-random/app/domain/entity"
)

type GoogleRepository interface {
	GetToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (*entity.User, error)
}
