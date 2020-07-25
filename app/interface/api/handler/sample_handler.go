package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ktakenaka/go-random/app/domain/service"
	"github.com/ktakenaka/go-random/app/interface/api/presenter"
	"github.com/ktakenaka/go-random/app/interface/persistence/mysql"
	"github.com/ktakenaka/go-random/app/usecase"
)

func getSamples(ctx *gin.Context) {
	sampleRepository := mysql.NewSampleRepository()
	sampleService := service.NewSampleService(sampleRepository)
	suCase := usecase.NewSampleUsecase(sampleRepository, sampleService)
	samples, err := suCase.ListSample()

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}

	sampleRes := make([]presenter.SampleResponse, 0)
	for _, sm := range samples {
		sampleRes = append(sampleRes, presenter.NewSampleResponse(sm))
	}
	ctx.JSON(http.StatusOK, gin.H{"samples": sampleRes})
}

func getSample(ctx *gin.Context) {
	sampleRepository := mysql.NewSampleRepository()
	sampleService := service.NewSampleService(sampleRepository)
	suCase := usecase.NewSampleUsecase(sampleRepository, sampleService)

	sample, err := suCase.FindSample(ctx.Params.ByName("id"))

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}
	sampleRes := presenter.NewSampleResponse(sample)
	ctx.JSON(http.StatusOK, gin.H{"sample": sampleRes})
}

func postSample(ctx *gin.Context) {
	sampleRepository := mysql.NewSampleRepository()
	sampleService := service.NewSampleService(sampleRepository)
	suCase := usecase.NewSampleUsecase(sampleRepository, sampleService)

	err := suCase.RegisterSample(ctx.PostForm("title"))
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func putSample(ctx *gin.Context) {
	sampleRepository := mysql.NewSampleRepository()
	sampleService := service.NewSampleService(sampleRepository)
	suCase := usecase.NewSampleUsecase(sampleRepository, sampleService)

	err := suCase.UpdateSample(ctx.Params.ByName("id"), ctx.PostForm("title"))
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func deleteSample(ctx *gin.Context) {
	sampleRepository := mysql.NewSampleRepository()
	sampleService := service.NewSampleService(sampleRepository)
	suCase := usecase.NewSampleUsecase(sampleRepository, sampleService)

	if err := suCase.DeleteSample(ctx.Params.ByName("id")); err != nil {
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
