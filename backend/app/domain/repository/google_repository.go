package repository

import (
	"context"

	"golang.org/x/oauth2"
)

type GoogleRepository interface {
	GetToken(ctx context.Context, code, nonce string) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, token *oauth2.Token) (map[string]interface{}, error)
}
