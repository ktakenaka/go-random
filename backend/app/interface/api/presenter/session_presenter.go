package presenter

// GoogleSessionRequest represents the request create session with Google
type GoogleSessionRequest struct {
	Code  string `json:"code" binding:"required"`
	Nonce string `json:"nonce" binding:"required"`
}

// SessionResponse represents the response of session creation
type SessionResponse struct {
	CSRFToken string `json:"csrf_token"`
}
