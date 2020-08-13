package jwtutil

import (
	"crypto/rand"
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
	csrfTokenLength = 32
	aExpDuration    = time.Hour * 12
	rExpDuration    = time.Hour * 24
)

// InitJWTSecret set jwtSecret
func InitJWTSecret(secret string) {
	once.Do(func() {
		jwtSecret = []byte(secret)
	})
}

// GenerateToken generates JWT for authentication
func GenerateToken(claims *AuthClaims) (aTokenStr, rTokenStr, csrfToken string, err error) {
	csrfToken, err = generateCSRFToken(csrfTokenLength)
	if err != nil {
		return "", "", "", err
	}

	claims.CSRFToken = csrfToken

	aTokenStr, err = generateJWT(claims, aExpDuration)
	if err != nil {
		return "", "", "", err
	}

	rTokenStr, err = generateJWT(claims, rExpDuration)

	return aTokenStr, rTokenStr, csrfToken, err
}

// TODO: consider using Redis to revoke token per user
func generateJWT(claims *AuthClaims, expDur time.Duration) (tokenStr string, err error) {
	issuedAt := time.Now()
	claims.StandardClaims = jwt.StandardClaims{
		IssuedAt:  issuedAt.Unix(),
		ExpiresAt: issuedAt.Add(expDur).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(jwtSecret)

	return tokenStr, err
}

func generateCSRFToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
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
		UserID: uint64(claims["user_id"].(float64)),
	}
	return authClaims, nil
}
