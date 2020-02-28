package bst_web_models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	OneTimePassword string `json:"otp,omitempty"`
}