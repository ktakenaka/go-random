package middleware

import (
	stdErrors "errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/helper/jwtutil"
)

const (
	accessKey      = "aa_jwt_access"
	refreshKey     = "aa_jwt_refresh"
	userIDKey      = "userID"
	csrfHeader     = "X-CSRF-TOKEN"
	cookieMaxAge   = 86400
	cookiePath     = "/api"
	cookieHTTPOnly = true
)

var (
	ignoreMethods  = []string{"GET", "HEAD", "OPTIONS"}
	cookieSameSite = judgeSameSiteMode(os.Getenv("ENV"))
	cookieSecure   = judegeSecure(os.Getenv("ENV"))
)

func judgeSameSiteMode(env string) http.SameSite {
	if env == "heroku" {
		// This configuration is for heroku review app
		// Cookie is not sent with cross domain unless setting None
		return http.SameSiteNoneMode
	}
	return http.SameSiteLaxMode
}

func judegeSecure(env string) bool {
	if env == "development" || env == "" {
		return false
	}
	return true
}

// CookieAuthenticator is for User Authentication and CSRF protection
type CookieAuthenticator struct{}

// NewCookieAuthenticator returns CookieAuthenticator
func NewCookieAuthenticator() CookieAuthenticator {
	return CookieAuthenticator{}
}

// AuthenticateAccess validates JWT and CSRF token, set userID and officeID
func (m *CookieAuthenticator) AuthenticateAccess(ctx *gin.Context) {
	m.authenticate(ctx, accessKey)
}

// AuthenticateRefresh validates JWT and CSRF token, set userID and officeID to refresh JWT
func (m *CookieAuthenticator) AuthenticateRefresh(ctx *gin.Context) {
	m.authenticate(ctx, refreshKey)
}

func (m *CookieAuthenticator) authenticate(ctx *gin.Context, jwtKey string) {
	var err error
	defer func() {
		if err != nil {
			// TODO: set error context
			ctx.Abort()
		}
	}()

	token, err := ctx.Cookie(jwtKey)
	if err != nil {
		return
	}

	claims, err := jwtutil.VerifyJWT(token)
	if err != nil {
		return
	}

	if !inArray(ctx.Request.Method, ignoreMethods) {
		csrfToken := ctx.Request.Header.Get(csrfHeader)
		if claims.CSRFToken != csrfToken {
			err = stdErrors.New("csrf detected")
			return
		}
	}

	ctx.Set(userIDKey, claims.UserID)
	ctx.Next()
}

func inArray(target string, arr []string) bool {
	for _, method := range arr {
		if target == method {
			return true
		}
	}
	return false
}

// SetAccessCookie is expected to to create & refresh JWT. i.e SessionHandler
func SetAccessCookie(ctx *gin.Context, token string) {
	ctx.SetSameSite(cookieSameSite)
	ctx.SetCookie(
		accessKey,
		token,
		cookieMaxAge,
		cookiePath,
		"",
		cookieSecure,
		cookieHTTPOnly,
	)
}

// SetRefreshCookie is expected to to create & refresh JWT. i.e SessionHandler
func SetRefreshCookie(ctx *gin.Context, token string) {
	ctx.SetSameSite(cookieSameSite)
	ctx.SetCookie(
		refreshKey,
		token,
		cookieMaxAge,
		cookiePath,
		"",
		cookieSecure,
		cookieHTTPOnly,
	)
}
