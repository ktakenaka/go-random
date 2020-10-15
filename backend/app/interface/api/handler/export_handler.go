package handler

import (
	"bytes"
	"encoding/csv"
	"io"
	"net/http"

	"github.com/gocarina/gocsv"
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
		b   *bytes.Buffer
		w   io.Writer
	)
	defer func() {
		if err != nil {
			hdl.SetError(ctx, err)
		}
	}()

	claims := hdl.JWTClaims(ctx)
	suCase := registry.InitializeSampleUsecase()
	samples, err := suCase.List(claims.UserID)
	if err != nil {
		return
	}

	if ctx.Query("charset") == "sjis" {
		w = transform.NewWriter(b, japanese.ShiftJIS.NewEncoder())
	} else {
		w = b
	}

	writer := csv.NewWriter(w)
	defer writer.Flush()

	if err = gocsv.MarshalCSV(samples, writer); err != nil {
		return
	}

	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=samples.csv")
	ctx.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	ctx.Writer.Header().Set("Expires", "0")
	ctx.Data(http.StatusOK, "text/csv", b.Bytes())
}
