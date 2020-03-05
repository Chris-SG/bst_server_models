package bst_models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	OneTimePassword string `json:"otp,omitempty"`
}

type EagateUser struct {
	Username string `json:"username"`
	Expired bool
}