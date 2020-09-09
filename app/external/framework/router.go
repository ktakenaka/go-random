package framework

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ktakenaka/go-random/app/interface/api/handler"
	"github.com/ktakenaka/go-random/app/interface/api/middleware"
)

func Handler() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	// TODO: refactor the condition.
	// It's not ideal to use `os.Getenv("ENV")` in several places
	if os.Getenv("ENV") == "" {
		router.Use(middleware.CorsMiddleware())
	}

	responseFmt := middleware.NewResponseFormatter()
	// TODO: use CustomRecovery after it bacomes available
	router.Use(responseFmt.Recovery)

	router.GET("/", root)

	v1NoAuth := router.Group("/api/v1")

	sampleHdl := handler.NewSampleHanlder()
	sample := v1NoAuth.Group("samples")
	sample.Use(responseFmt.Format)
	sample.GET("", sampleHdl.Index)
	sample.GET("/:id", sampleHdl.Show)
	sample.POST("", sampleHdl.Create)
	sample.PUT("/:id", sampleHdl.Update)
	sample.DELETE("/:id", sampleHdl.Delete)

	sessionHdl := handler.NewSessionHandler()
	session := v1NoAuth.Group("/sessions")
	session.POST("/google", sessionHdl.CreateWithGoogle)

	v1Auth := router.Group("/api/v1")
	cookieAuth := middleware.NewCookieAuthenticator()
	v1Auth.Use(cookieAuth.AuthenticateAccess)

	return router
}

func root(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "root"})
}
