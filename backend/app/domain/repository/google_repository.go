package repository

import (
	"golang.org/x/oauth2"
)

type GoogleRepository interface {
	GetToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (map[string]interface{}, error)
}
