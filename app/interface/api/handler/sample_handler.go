package handler


import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSmaple(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	fmt.Println(id)
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func AddSampleHanlder(g *gin.RouterGroup) {
	g.GET("/:id", getSmaple)
}
