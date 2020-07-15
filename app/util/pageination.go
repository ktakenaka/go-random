package util

import (
	"github.com/Unknown/com"
	"github.com/gin-gonic/gin"

	"app/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	// pageinationのpage番号か
	page, _ := com.StrTo(c.Query("page")).Int()
	
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
