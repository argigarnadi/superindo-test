package model

import (
	"github.com/google/uuid"
)

type RequestRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseRegister struct {
	UserId   uuid.UUID `json:"userId"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	CreateAt string    `json:"createAt"`
}
type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
