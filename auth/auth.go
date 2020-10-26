package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	Repo *AuthRepo
}

func NewService() *AuthService {
	return &AuthService{
		Repo: NewAuthRepo(),
	}
}

type Token struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type Claims struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type Login struct {
	UserID   int    `json:"userID"`
	Email    string `json:"email"`
	Password string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
