package domain

import "context"

type User struct {
	Username  string
	FullName  string
	Email     string
	Password  string
	UpdatedAt string
}

type UserUseCase interface {
	GetAll(ctx context.Context) ([]User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
}
