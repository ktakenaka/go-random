package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

const (
	sessionStore = "session"
	secretKey    = "mysecret"
)

func NewCSRFStore() gin.HandlerFunc {
	store := cookie.NewStore([]byte("secret"))
	return sessions.Sessions(sessionStore, store)
}

func NewGinCSRFMiddleware() gin.HandlerFunc {
	return csrf.Middleware(
		csrf.Options{
			Secret:    secretKey,
			ErrorFunc: errorFunc,
		},
	)
}

func errorFunc(ctx *gin.Context) {
	ctx.String(http.StatusBadRequest, "csrf token mismatch")
	ctx.Abort()
}

func GetCSRFToken(ctx *gin.Context) {
	ctx.String(http.StatusOK, csrf.GetToken(ctx))
}
