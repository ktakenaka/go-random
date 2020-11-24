package handler

import (
	"encoding/csv"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/jinzhu/copier"

	"github.com/ktakenaka/go-random/backend/app/interface/api/presenter"
	"github.com/ktakenaka/go-random/backend/app/registry"
	"github.com/ktakenaka/go-random/backend/app/usecase/dto"
	"github.com/ktakenaka/go-random/backend/pkg/logger"
)

// SampleHandler is the sample
type SampleHandler struct {
	BaseHandler
}

// NewSampleHandler is a constructor for Samplehandler
func NewSampleHandler() *SampleHandler {
	return &SampleHandler{}
}

// Index returns the list of samples
func (hdl *SampleHandler) Index(ctx *gin.Context) {
	var err error
	defer func() {
		hdl.SetError(ctx, err)
	}()

	q, err := dto.NewJSONAPIQueryFromContext(ctx)
	if err != nil {
		return
	}

	claims := hdl.JWTClaims(ctx)

	suCase := registry.InitializeSampleUsecase()
	samples, err := suCase.List(claims.UserID, q)
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

// Show returns a sample
func (hdl *SampleHandler) Show(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		if err != nil {
			logger.Error(err)
		}
		hdl.SetError(ctx, err)
	}()

	id, err := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		return
	}

	claims := hdl.JWTClaims(ctx)

	sample, err := suCase.Find(claims.UserID, id)
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

// Create creates a sample
func (hdl *SampleHandler) Create(ctx *gin.Context) {
	var (
		err    error
		dtoReq dto.CreateSample
	)
	defer func() {
		hdl.SetError(ctx, err)
	}()

	if err = ctx.ShouldBindJSON(&dtoReq); err != nil {
		return
	}

	claims := hdl.JWTClaims(ctx)
	dtoReq.UserID = claims.UserID

	suCase := registry.InitializeSampleUsecase()
	if err = suCase.Create(dtoReq); err != nil {
		return
	}

	hdl.SetData(ctx, "OK")
	hdl.SetMeta(ctx, presenter.CodeCreated)
}

// Update updates a sample
func (hdl *SampleHandler) Update(ctx *gin.Context) {
	var (
		err    error
		dtoReq dto.UpdateSample
	)
	defer func() {
		hdl.SetError(ctx, err)
	}()

	if err = ctx.ShouldBindJSON(&dtoReq); err != nil {
		return
	}

	id, err := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		return
	}

	dtoReq.ID = id

	claims := hdl.JWTClaims(ctx)
	dtoReq.UserID = claims.UserID

	suCase := registry.InitializeSampleUsecase()
	if err = suCase.Update(dtoReq); err != nil {
		return
	}

	hdl.SetData(ctx, "ok")
	hdl.SetMeta(ctx, presenter.CodeSuccess)
}

// Delete deletes a sample
func (hdl *SampleHandler) Delete(ctx *gin.Context) {
	suCase := registry.InitializeSampleUsecase()

	var err error
	defer func() {
		hdl.SetError(ctx, err)
	}()

	id, err := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		return
	}

	claims := hdl.JWTClaims(ctx)

	if err = suCase.Delete(claims.UserID, id); err != nil {
		return
	}

	hdl.SetData(ctx, "ok")
	hdl.SetMeta(ctx, presenter.CodeSuccess)
}

// Import csv
func (hdl *SampleHandler) Import(ctx *gin.Context) {
	var (
		err error
	)
	defer func() {
		hdl.SetError(ctx, err)
	}()

	header, err := ctx.FormFile("file")
	if err != nil {
		return
	}

	file, err := header.Open()
	if err != nil {
		return
	}

	prsSamples := []presenter.SampleCSVPresenter{}
	reader := csv.NewReader(file)

	if err = gocsv.UnmarshalCSV(reader, &prsSamples); err != nil {
		return
	}

	var dtoSamples []dto.ImportSample
	if err = copier.Copy(&dtoSamples, &prsSamples); err != nil {
		return
	}

	suCase := registry.InitializeSampleUsecase()
	if err = suCase.Import(dtoSamples); err != nil {
		return
	}

	hdl.SetData(ctx, "ok")
	hdl.SetMeta(ctx, presenter.CodeSuccess)
}
