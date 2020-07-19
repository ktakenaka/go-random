package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ktakenaka/go-random/app/interface/api/handler"
)

func Handler() *gin.Engine {
	router := gin.Default()

	router.GET("/", root)

	v1 := router.Group("/api/v1")
	handler.AddSampleHanlder(v1.Group("/samples"))

	return router
}

func root(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "root"})
}
