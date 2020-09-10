package middleware

import (
	stdErrors "errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/interface/api/presenter"
	"github.com/ktakenaka/go-random/helper/jwtutil"
)

const (
	accessKey      = "aa_jwt_access"
	refreshKey     = "aa_jwt_refresh"
	jwtClaimsKey   = "jwtClaims"
	csrfHeader     = "X-CSRF-TOKEN"
	cookieMaxAge   = 86400
	cookiePath     = "/api"
	cookieHTTPOnly = true
)

var (
	ignoreMethods  []string      = []string{"GET", "HEAD", "OPTIONS"}
	cookieSameSite http.SameSite = http.SameSiteLaxMode
	cookieSecure   bool          = true
)

// InitJWTCookieOpt initializes Cookie option for JWT
func InitJWTCookieOpt(env string) {
	if env == "heroku" {
		cookieSameSite = http.SameSiteNoneMode
	}

	if env == "development" {
		cookieSecure = false
	}
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
			SetErrorResponse(ctx, err)
			ctx.Abort()
		}
	}()

	token, err := ctx.Cookie(jwtKey)
	if err != nil {
		SetMetaResponse(ctx, presenter.CodeUnauthorized)
		return
	}

	claims, err := jwtutil.VerifyJWT(token)
	if err != nil {
		SetMetaResponse(ctx, presenter.CodeUnauthorized)
		return
	}

	if !inArray(ctx.Request.Method, ignoreMethods) {
		csrfToken := ctx.Request.Header.Get(csrfHeader)
		if claims.CSRFToken != csrfToken {
			err = stdErrors.New("csrf detected")
			SetMetaResponse(ctx, presenter.CodeForbidden)
			return
		}
	}

	jwtClaims := JWTClaims{
		UserID: claims.UserID,
	}
	ctx.Set(jwtClaimsKey, jwtClaims)
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

// ExtractClaims extracts AuthClaims from context
func ExtractClaims(ctx *gin.Context) (JWTClaims, error) {
	claims, ok := ctx.Get(jwtClaimsKey)
	if !ok {
		// We will never have this error because AuthenticateAccess raises error before accessing here.
		// It means jwtClaims always exists that AuthenticateAccess doesn't raise error.
		return JWTClaims{}, stdErrors.New("failed to extract jwtClaims")
	}

	return claims.(JWTClaims), nil
}

// JWTClaims is a data transfer object
type JWTClaims struct {
	UserID uint64
}
