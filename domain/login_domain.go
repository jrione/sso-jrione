package domain

import "context"

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginUseCase interface {
	CheckUser(c context.Context, usecase string) (*User, error)
}
