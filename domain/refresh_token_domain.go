package domain

import "context"

type RefreshTokenData struct {
	Username     string
	RefreshToken string
}

type RefreshTokenRepository interface {
	StoreRefreshToken(context.Context, RefreshTokenData) (err error)
	GetRefreshToken(context.Context, string) (*RefreshTokenData, error)
	// UpdateRefreshToken(context.Context, string) (string, error)
	DeleteRefreshToken(context.Context, string) (bool, error)
}

type RefreshTokenUseCase interface {
	StoreRefreshToken(context.Context, RefreshTokenData) (err error)
	GetRefreshToken(context.Context, string) (*RefreshTokenData, error)
	// UpdateRefreshToken(context.Context, string) (string, error)
	DeleteRefreshToken(context.Context, string) (bool, error)
}
