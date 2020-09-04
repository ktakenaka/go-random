package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	localhost = "http://127.0.0.1:3000"
)

// CorsMiddleware is for development where FE and BE use different domain
func CorsMiddleware() gin.HandlerFunc {
	middleware := cors.New(cors.Config{
		AllowOrigins:     []string{localhost},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "Origin", "Cookie", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowCredentials: true,
	})

	return middleware
}
