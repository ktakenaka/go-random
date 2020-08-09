package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/config"
	"github.com/ktakenaka/go-random/app/registry"
)

func AddSessionHandler(g *gin.RouterGroup) {
	g.POST("/google", createSessionWithGoogle)
}

func createSessionWithGoogle(ctx *gin.Context) {
	uc := registry.InitializeSignInUsecase()
	user, err := uc.Execute(ctx.PostForm("code"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	// investigate JWT for authentication, including refresh
	csrfToken := "csrf_token"
	expDuration := time.Hour * 24

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(expDuration).Unix()
	claims["user_id"] = user.ID //TODO: make user_id hash
	claims["csrf_token"] = csrfToken
	tokenString, _ := token.SignedString([]byte(config.GetJWTSecret()))

	ctx.SetCookie(
		"aa_jwt_access",
		tokenString,
		86400, // 1 day
		"/api",
		"",
		false, // convert it to true unless it's development
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{"data": csrfToken})
}
