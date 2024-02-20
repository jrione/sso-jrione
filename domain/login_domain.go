package domain

import (
	"context"

	"github.com/golang-jwt/jwt"
)

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginUseCase interface {
	CheckUser(c context.Context, usecase string) (*User, error)
	CreateAccessToken(user *User, secret string, expire int) (string, error)
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
