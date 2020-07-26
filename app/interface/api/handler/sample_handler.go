package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/ktakenaka/go-random/app/domain/service"
	"github.com/ktakenaka/go-random/app/external/database"
	"github.com/ktakenaka/go-random/app/interface/api/presenter"
	"github.com/ktakenaka/go-random/app/interface/persistence/mysql"
	"github.com/ktakenaka/go-random/app/usecase"
)

func AddSampleHanlder(g *gin.RouterGroup) {
	g.GET("", getSamples)
	g.GET("/:id", getSample)
	g.POST("", postSample)
	g.PUT("/:id", putSample)
	g.DELETE("/:id", deleteSample)
}

func getSamples(ctx *gin.Context) {
	db := database.MySQLConnection()
	sampleRepository := mysql.NewSampleRepository(db)
	sampleService := service.NewSampleService(sampleRepository)
	suCase := usecase.NewSampleUsecase(sampleRepository, sampleService)
	samples, err := suCase.ListSample()

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}

	sampleRes := make([]presenter.SampleResponse, 0)
	if err := copier.Copy(&sampleRes, &samples); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}
	ctx.JSON(http.StatusOK, gin.H{"samples": sampleRes})
}

func getSample(ctx *gin.Context) {
	db := database.MySQLConnection()
	sampleRepository := mysql.NewSampleRepository(db)
	sampleService := service.NewSampleService(sampleRepository)
	suCase := usecase.NewSampleUsecase(sampleRepository, sampleService)

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id must be integer"})
	}

	sample, err := suCase.FindSample(id)

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}

	var sampleRes presenter.SampleResponse
	if err := copier.Copy(&sampleRes, &sample); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}
	ctx.JSON(http.StatusOK, gin.H{"sample": sampleRes})
}

func postSample(ctx *gin.Context) {
	db := database.MySQLConnection()
	sampleRepository := mysql.NewSampleRepository(db)
	sampleService := service.NewSampleService(sampleRepository)
	suCase := usecase.NewSampleUsecase(sampleRepository, sampleService)

	if err := suCase.RegisterSample(ctx.PostForm("title")); err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func putSample(ctx *gin.Context) {
	db := database.MySQLConnection()
	sampleRepository := mysql.NewSampleRepository(db)
	sampleService := service.NewSampleService(sampleRepository)
	suCase := usecase.NewSampleUsecase(sampleRepository, sampleService)

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id must be integer"})
	}

	if err := suCase.UpdateSample(id, ctx.PostForm("title")); err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func deleteSample(ctx *gin.Context) {
	db := database.MySQLConnection()
	sampleRepository := mysql.NewSampleRepository(db)
	sampleService := service.NewSampleService(sampleRepository)
	suCase := usecase.NewSampleUsecase(sampleRepository, sampleService)

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id must be integer"})
	}

	if err := suCase.DeleteSample(id); err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
