package restclient

import (
	"context"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"

	"github.com/ktakenaka/go-random/app/domain/entity"
)

type GoogleRepository struct {
	oauth2Cnf *oauth2.Config
}

func NewGoogleRepository(cnf *oauth2.Config) *GoogleRepository {
	return &GoogleRepository{
		oauth2Cnf: cnf,
	}
}

func (r *GoogleRepository) GetToken(code string) (*oauth2.Token, error) {
	oauth2Tkn, err := r.oauth2Cnf.Exchange(
		context.Background(),
		code,
	)
	if err != nil {
		return nil, err
	}
	return oauth2Tkn, nil
}

func (r *GoogleRepository) GetUserInfo(token *oauth2.Token) (*entity.User, error) {
	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		return nil, err
	}

	userInfo, err := provider.UserInfo(context.Background(), oauth2.StaticTokenSource(token))
	if err != nil {
		return nil, err
	}

	body := make(map[string]interface{})
	err = userInfo.Claims(&body)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		GoogleSub: body["sub"].(string),
		Email:     body["email"].(string),
	}

	return &user, nil
}
