package jwtutil

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthClaims the claim for authentication
type AuthClaims struct {
	jwt.StandardClaims
	UserID    string    `json:"user_id,omitempty"`
	CSRFToken string    `json:"csrf_token,omitempty"`
	IssueTime time.Time `json:"-"`
}
