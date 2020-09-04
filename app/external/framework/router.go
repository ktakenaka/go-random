package framework

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ktakenaka/go-random/app/interface/api/handler"
	"github.com/ktakenaka/go-random/app/interface/api/middleware"
)

func Handler() *gin.Engine {
	router := gin.Default()

	// TODO: refactor the condition.
	// It's not ideal to use `os.Getenv("ENV")` in several places
	if os.Getenv("ENV") == "" {
		router.Use(middleware.CorsMiddleware())
	}

	router.GET("/", root)

	v1NoAuth := router.Group("/api/v1")
	handler.AddSessionHandler(v1NoAuth.Group("/sessions"))
	handler.AddSampleHanlder(v1NoAuth.Group("/samples"))

	v1Auth := router.Group("/api/v1")
	cookieAuth := middleware.NewCookieAuthenticator()
	v1Auth.Use(cookieAuth.AuthenticateAccess)
	// When we want to test Auth, remove comment out
	// handler.AddSampleHanlder(v1Auth.Group("/samples"))

	// These are trial purpose, not for use
	if os.Getenv("ENV") == "development" {
		authMiddleware := middleware.NewGinJWTMiddleware()
		router.POST("/login", authMiddleware.LoginHandler)

		auth := router.Group("/auth")
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.Use(authMiddleware.MiddlewareFunc())
		auth.GET("/hello", middleware.HelloHandler)

		csrfTrial := router.Group("/csrf")
		csrfTrial.Use(middleware.NewCSRFStore())
		csrfTrial.Use(middleware.NewGinCSRFMiddleware())
		csrfTrial.GET("/protected", middleware.GetCSRFToken)
		csrfTrial.POST("/protected", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
	}

	return router
}

func root(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "root"})
}
