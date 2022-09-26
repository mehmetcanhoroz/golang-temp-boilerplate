package models

type AuthRegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthSuccessfulResponse struct {
	Token string `json:"token"`
}
