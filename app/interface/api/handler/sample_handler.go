package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/ktakenaka/go-random/app/interface/api/presenter"
	"github.com/ktakenaka/go-random/app/registry"
)

type SampleHanlder struct{}

func NewSampleHanlder() *SampleHanlder {
	return &SampleHanlder{}
}

func (hdl *SampleHanlder) Index(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()
	samples, err := suCase.ListSample()

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}

	sampleRes := make([]presenter.SampleResponse, 0)
	if err := copier.Copy(&sampleRes, &samples); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": sampleRes})
}

func (hdl *SampleHanlder) Show(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "id must be integer"})
	}

	sample, err := suCase.FindSample(id)

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "NG"})
	}

	var sampleRes presenter.SampleResponse
	if err := copier.Copy(&sampleRes, &sample); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "NG"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": sampleRes})
}

func (hdl *SampleHanlder) Create(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var req presenter.SampleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "param is invalid"})
	}

	if err := suCase.RegisterSample(req.Title); err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (hdl *SampleHanlder) Update(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var req presenter.SampleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "param is invalid"})
	}

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id must be integer"})
		return
	}

	if err := suCase.UpdateSample(id, req.Title); err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "NG"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (hdl *SampleHanlder) Delete(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

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
