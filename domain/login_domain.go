package domain

import (
	"context"

	"github.com/golang-jwt/jwt"
)

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

type LogoutRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginUseCase interface {
	CheckUser(c context.Context, usecase string) (*User, error)
	CreateAccessToken(user *User, secret string, expire int) (string, error)
	CreateRefreshToken(user *User, secret string, expire int) (string, error)
	GetUsernameFromClaim(user string, secret string) (string, error)
}
