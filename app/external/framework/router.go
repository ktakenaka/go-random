package framework

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ktakenaka/go-random/app/interface/api/handler"
	"github.com/ktakenaka/go-random/app/interface/api/middleware"
)

func Handler() *gin.Engine {
	router := gin.Default()

	router.GET("/", root)

	v1 := router.Group("/api/v1")
	handler.AddSampleHanlder(v1.Group("/samples"))
	handler.AddSessionHandler(v1.Group("/sessions"))

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

	return router
}

func root(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "root"})
}
