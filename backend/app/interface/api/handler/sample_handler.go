package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/ktakenaka/go-random/app/interface/api/presenter"
	"github.com/ktakenaka/go-random/app/registry"
)

// SampleHandler is the sample
type SampleHandler struct {
	BaseHandler
}

func NewSampleHandler() *SampleHandler {
	return &SampleHandler{}
}

func (hdl *SampleHandler) Index(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		hdl.SetError(ctx, err)
	}()

	samples, err := suCase.ListSample()

	if err != nil {
		return
	}

	sampleRes := make([]presenter.SampleResponse, 0)
	if err = copier.Copy(&sampleRes, &samples); err != nil {
		return
	}

	hdl.SetData(ctx, sampleRes)
	hdl.SetMeta(ctx, presenter.CodeSuccess)
}

func (hdl *SampleHandler) Show(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		hdl.SetError(ctx, err)
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

	hdl.SetData(ctx, sampleRes)
	hdl.SetMeta(ctx, presenter.CodeSuccess)
}

func (hdl *SampleHandler) Create(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		hdl.SetError(ctx, err)
	}()

	var req presenter.SampleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	if err = suCase.RegisterSample(req.Title); err != nil {
		return
	}

	hdl.SetData(ctx, "OK")
	hdl.SetMeta(ctx, presenter.CodeCreated)
}

func (hdl *SampleHandler) Update(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		hdl.SetError(ctx, err)
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

	hdl.SetData(ctx, "ok")
	hdl.SetMeta(ctx, presenter.CodeSuccess)
}

func (hdl *SampleHandler) Delete(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		hdl.SetError(ctx, err)
	}()

	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		return
	}

	if err = suCase.DeleteSample(id); err != nil {
		return
	}

	hdl.SetData(ctx, "ok")
	hdl.SetMeta(ctx, presenter.CodeSuccess)
}
