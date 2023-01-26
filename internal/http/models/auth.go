package models

// RegisterInput is the http register user model
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginInput is the http login user model
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TokenResponse is the response model for a token response
type TokenResponse struct {
	Token string `json:"token,omitempty"`
}
