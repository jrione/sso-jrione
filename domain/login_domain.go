package domain

import "context"

type LoginRequest struct {
	Username string
	Password string
}

type LoginUseCase interface {
	CheckUser(c context.Context, usecase string) (*User, error)
}
