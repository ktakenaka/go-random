package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ktakenaka/go-random/app/domain/entity"
)

func getSamples(ctx *gin.Context) {
	samples, err := entity.ListSamples()
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"samples": samples})
	}
}

func getSample(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	sample, err := entity.FindSample(id)

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"sample": sample})
	}
}

func postSample(ctx *gin.Context) {
	title := ctx.PostForm("title")
	err := entity.CreateSample(title)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func putSample(ctx *gin.Context) {
	sample, err := entity.FindSample(ctx.Params.ByName("id"))
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}

	sample.Title = ctx.PostForm("title")
	err = sample.Save()

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func deleteSample(ctx *gin.Context) {
	sample, err := entity.FindSample(ctx.Params.ByName("id"))
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}

	err = sample.Delete()

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func AddSampleHanlder(g *gin.RouterGroup) {
	g.GET("", getSamples)
	g.GET("/:id", getSample)
	g.POST("", postSample)
	g.PUT("/:id", putSample)
	g.DELETE("/:id", deleteSample)
}
