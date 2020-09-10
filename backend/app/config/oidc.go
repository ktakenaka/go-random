package config

import (
	"sync"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	once           sync.Once
	googleOauthCnf *oauth2.Config
)

func InitGoogleOIDCCnf(redirectURL, clientID, clientSecret string) {
	once.Do(func() {
		googleOauthCnf = &oauth2.Config{
			RedirectURL:  redirectURL,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
			Endpoint:     google.Endpoint,
		}
	})
}

func GetGoogleOauthConfig() *oauth2.Config {
	return googleOauthCnf
}
