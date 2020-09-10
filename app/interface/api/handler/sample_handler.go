package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/ktakenaka/go-random/app/interface/api/middleware"
	"github.com/ktakenaka/go-random/app/interface/api/presenter"
	"github.com/ktakenaka/go-random/app/registry"
)

type SampleHanlder struct{}

func NewSampleHanlder() *SampleHanlder {
	return &SampleHanlder{}
}

func (hdl *SampleHanlder) Index(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		middleware.SetError(ctx, err)
	}()

	samples, err := suCase.ListSample()

	if err != nil {
		return
	}

	sampleRes := make([]presenter.SampleResponse, 0)
	if err = copier.Copy(&sampleRes, &samples); err != nil {
		return
	}

	middleware.SetDataResponse(ctx, sampleRes)
}

func (hdl *SampleHanlder) Show(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		middleware.SetError(ctx, err)
	}()

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		return
	}

	sample, err := suCase.FindSample(id)
	if err != nil {
		return
	}

	var sampleRes presenter.SampleResponse
	if err = copier.Copy(&sampleRes, &sample); err != nil {
		return
	}

	middleware.SetDataResponse(ctx, sampleRes)
}

func (hdl *SampleHanlder) Create(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		middleware.SetError(ctx, err)
	}()

	var req presenter.SampleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	if err = suCase.RegisterSample(req.Title); err != nil {
		return
	}

	middleware.SetDataResponse(ctx, "OK")
}

func (hdl *SampleHanlder) Update(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		middleware.SetError(ctx, err)
	}()

	var req presenter.SampleRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		return
	}

	if err = suCase.UpdateSample(id, req.Title); err != nil {
		return
	}

	middleware.SetDataResponse(ctx, "ok")
}

func (hdl *SampleHanlder) Delete(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		middleware.SetError(ctx, err)
	}()

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		return
	}

	if err = suCase.DeleteSample(id); err != nil {
		return
	}

	middleware.SetDataResponse(ctx, "ok")
}
