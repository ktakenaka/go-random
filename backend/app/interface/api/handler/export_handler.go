package handler

import (
	"encoding/csv"
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/gin-gonic/gin"
	"github.com/ktakenaka/go-random/backend/app/registry"
)

// ExportHandler is the sample
type ExportHandler struct {
	BaseHandler
}

// NewExportHandler is a constructor for ExportHandler
func NewExportHandler() *ExportHandler {
	return &ExportHandler{}
}

// SamplesExport exports samples
func (hdl *ExportHandler) SamplesExport(ctx *gin.Context) {
	var (
		err error
		w   io.Writer
	)
	defer func() {
		if err != nil {
			hdl.SetError(ctx, err)
		}
	}()

	claims := hdl.JWTClaims(ctx)

	if ctx.Query("charset") == "sjis" {
		w = transform.NewWriter(ctx.Writer, japanese.ShiftJIS.NewEncoder())
	} else {
		w = ctx.Writer
	}

	suCase := registry.InitializeSampleUsecase()
	samples, err := suCase.List(claims.UserID)
	if err != nil {
		return
	}

	writer := csv.NewWriter(w)
	defer writer.Flush()

	data := make([]string, len(samples))
	for _, s := range samples {
		d := []string{s.Title, s.Content.String}
		data = append(data, d...)
	}
	if err = writer.Write(data); err != nil {
		return
	}

	ctx.Writer.Header().Set("Content-Type", "text/csv")
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=samples.csv")
	ctx.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	ctx.Writer.Header().Set("Expires", "0")
}
