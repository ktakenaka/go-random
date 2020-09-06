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

// NewCSRFStore is the trial implementation.
// The usage is as below
/*
csrfTrial := router.Group("/csrf")
csrfTrial.Use(middleware.NewCSRFStore())
csrfTrial.Use(middleware.NewGinCSRFMiddleware())
csrfTrial.GET("/protected", middleware.GetCSRFToken)
csrfTrial.POST("/protected", func(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
})
*/

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
