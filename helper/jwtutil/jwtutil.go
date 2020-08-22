package jwtutil

import (
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	once      sync.Once
	jwtSecret []byte
)

// Error constant for jwt auth
var (
	ErrInvalidAlg = fmt.Errorf("invalid alg")
)

const (
	aExpDuration = time.Hour * 12
	rExpDuration = time.Hour * 24
)

// InitJWTSecret set jwtSecret
func InitJWTSecret(secret string) {
	once.Do(func() {
		jwtSecret = []byte(secret)
	})
}

// GenerateToken generates JWT for authentication
func GenerateToken(claims *AuthClaims) (aTokenStr, rTokenStr string, err error) {
	aTokenStr, err = generateJWT(claims, aExpDuration)
	if err != nil {
		return "", "", err
	}

	rTokenStr, err = generateJWT(claims, rExpDuration)

	return aTokenStr, rTokenStr, err
}

// TODO: consider using Redis to revoke token per user
func generateJWT(claims *AuthClaims, expDur time.Duration) (tokenStr string, err error) {
	claims.StandardClaims = jwt.StandardClaims{
		IssuedAt:  claims.IssueTime.Unix(),
		ExpiresAt: claims.IssueTime.Add(expDur).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(jwtSecret)

	return tokenStr, err
}

// VerifyJWT validates JWT and extract userId and officeID
func VerifyJWT(tokenStr string) (AuthClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidAlg
		}
		return jwtSecret, nil
	})

	if err != nil {
		return AuthClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return AuthClaims{}, err
	}

	authClaims := AuthClaims{
		UserID:    uint64(claims["user_id"].(float64)),
		CSRFToken: claims["csrf_token"].(string),
	}
	return authClaims, nil
}
