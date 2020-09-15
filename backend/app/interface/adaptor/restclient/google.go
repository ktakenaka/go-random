package restclient

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/ktakenaka/go-random/backend/app/config"
)

const (
	idTokenAttr = "id_token"
)

// GoogleRepository follows OIDC
// https://developers.google.com/identity/protocols/oauth2/openid-connect
type GoogleRepository struct {
	oauth2Cnf *oauth2.Config
}

// NewGoogleRepository returns GoogleRepository
func NewGoogleRepository(cnf *oauth2.Config) *GoogleRepository {
	return &GoogleRepository{
		oauth2Cnf: cnf,
	}
}

// GetToken calls Google API to get token and verify id_token
func (r *GoogleRepository) GetToken(ctx context.Context, code, nonce string) (*oauth2.Token, error) {
	oauth2Tkn, err := r.oauth2Cnf.Exchange(
		ctx,
		code,
	)
	if err != nil {
		return nil, err
	}

	rawIDTkn, ok := oauth2Tkn.Extra(idTokenAttr).(string)
	if !ok {
		// TODO: error handling
		return nil, fmt.Errorf("failed to get id_token")
	}

	verifier := config.GetGoogleVerifier(ctx)
	idToken, err := verifier.Verify(ctx, rawIDTkn)
	if err != nil {
		// TODO: error handling
		return nil, fmt.Errorf("invalid id_token")
	}

	if idToken.Nonce != nonce {
		// TODO: error handling
		return nil, fmt.Errorf("wrong nonce")
	}

	return oauth2Tkn, nil
}

// GetUserInfo gets user information from Google
func (r *GoogleRepository) GetUserInfo(ctx context.Context, token *oauth2.Token) (map[string]interface{}, error) {
	provider, err := config.GetGoogleProvider(ctx)
	if err != nil {
		return nil, err
	}

	userInfo, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(token))
	if err != nil {
		return nil, err
	}

	body := make(map[string]interface{})
	err = userInfo.Claims(&body)
	if err != nil {
		return nil, err
	}

	// TODO: define the struct to manage interface
	return body, nil
}
