package jwtutil

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// AuthClaims the claim for authentication
type AuthClaims struct {
	jwt.StandardClaims
	UserID    uint64    `json:"user_id,omitempty"`
	CSRFToken string    `json:"csrf_token,omitempty"`
	IssueTime time.Time `json:"-"`
}
