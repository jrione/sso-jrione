package usecase

import (
	"context"
	"time"

	"github.com/jrione/sso-jrione/domain"
)

type refreshTokenUseCase struct {
	refreshTokenRepository domain.RefreshTokenRepository
	contextTimeout         time.Duration
}

func NewRefreshTokenUseCase(refreshTokenRepo domain.RefreshTokenRepository, timeout time.Duration) domain.RefreshTokenUseCase {
	return &refreshTokenUseCase{
		refreshTokenRepository: refreshTokenRepo,
		contextTimeout:         timeout,
	}
}

func (r *refreshTokenUseCase) StoreRefreshToken(ctx context.Context, refreshTokenData domain.RefreshTokenData) (err error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	err = r.refreshTokenRepository.StoreRefreshToken(ctx, refreshTokenData)
	return
}

func (r *refreshTokenUseCase) GetRefreshToken(ctx context.Context, username string) (res *domain.RefreshTokenData, err error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	res, err = r.refreshTokenRepository.GetRefreshToken(ctx, username)
	return
}

// func (r *refreshTokenUseCase) UpdateRefreshToken(ctx context.Context, refreshToken string) (string, error) {
// 	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
// 	defer cancel()
// 	return "", nil
// }

// func (r *refreshTokenUseCase) DeleteRefreshToken(ctx context.Context, refreshToken string) (ok bool, err error) {
// 	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
// 	defer cancel()

// 	return false, nil
// }
