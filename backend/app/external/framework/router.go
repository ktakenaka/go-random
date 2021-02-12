package framework

import (
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ktakenaka/go-random/backend/app/interface/api/handler"
	"github.com/ktakenaka/go-random/backend/app/interface/api/middleware"
)

// Handler handle
func Handler(pprofEnabled bool) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	// TODO: refactor the condition.
	// It's not ideal to use `os.Getenv("ENV")` in several places
	if os.Getenv("ENV") == "development" {
		router.Use(middleware.CorsMiddleware())
	}

	responseFmt := middleware.NewResponseFormatter()
	// TODO: use CustomRecovery after it bacomes available
	router.Use(responseFmt.PanicRecovery)
	router.Use(responseFmt.Format)

	router.GET("/", root)

	v1NoAuth := router.Group("/api/v1")

	sessionHdl := handler.NewSessionHandler()
	session := v1NoAuth.Group("/sessions")
	session.POST("/google", sessionHdl.CreateWithGoogle)

	v1Auth := router.Group("/api/v1")
	jwtAuth := middleware.NewJWTAuthenticator()
	v1Auth.Use(jwtAuth.AuthenticateAccess)

	sampleHdl := handler.NewSampleHandler()
	sample := v1Auth.Group("samples")
	sample.GET("", sampleHdl.Index)
	sample.GET("/:id", sampleHdl.Show)
	sample.POST("", sampleHdl.Create)
	sample.PUT("/:id", sampleHdl.Update)
	sample.DELETE("/:id", sampleHdl.Delete)
	sample.POST("/import", sampleHdl.Import)

	exportHdl := handler.NewExportHandler()
	export := v1Auth.Group("export")
	export.GET("/samples", exportHdl.SamplesExport)

	if pprofEnabled {
		pprf := router.Group("pprof")
		pprf.GET("", pprofHandler(pprof.Index))
		pprf.GET("/cmdline", pprofHandler(pprof.Cmdline))
		pprf.GET("/profile", pprofHandler(pprof.Profile))
		pprf.POST("/symbol", pprofHandler(pprof.Symbol))
		pprf.GET("/symbol", pprofHandler(pprof.Symbol))
		pprf.GET("/trace", pprofHandler(pprof.Trace))
		pprf.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		pprf.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		pprf.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		pprf.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		pprf.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		pprf.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}

	return router
}

func root(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "root"})
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
