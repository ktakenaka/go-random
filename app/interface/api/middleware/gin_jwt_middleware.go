package middleware

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/registry"
)

type JWTPayload struct {
	UserID uint64
}

const (
	secret      = "secret-for-dev"
	identityKey = "user_id"
)

func NewGinJWTMiddleware() *jwt.GinJWTMiddleware {
	googleAuthMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "go-random",
		Key:         []byte(secret),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour * 24, // TODO: check exp
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(JWTPayload); ok {
				return jwt.MapClaims{
					identityKey: v.UserID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx *gin.Context) interface{} {
			claims := jwt.ExtractClaims(ctx)
			return &JWTPayload{
				UserID: uint64(claims[identityKey].(float64)),
			}
		},
		Authenticator: func(ctx *gin.Context) (interface{}, error) {
			uc := registry.InitializeSignInUsecase()
			user, err := uc.Execute(ctx.PostForm("code"))
			if err != nil {
				return JWTPayload{}, err
			}
			return JWTPayload{UserID: user.ID}, nil
		},
		Authorizator: func(data interface{}, ctx *gin.Context) bool {
			return true
		},
		Unauthorized: func(ctx *gin.Context, code int, message string) {
			ctx.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:    "cookie: go-random_jwt",
		SendCookie:     true,
		CookieName:     "go-random_jwt",
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteLaxMode,
		TimeFunc:       time.Now,
	})
	if err != nil {
		log.Fatal("jwt error" + err.Error())
	}
	return googleAuthMiddleware
}

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(http.StatusOK, gin.H{"message": claims})
}
