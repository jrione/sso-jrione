package usecase

import (
	"context"
	"time"

	"github.com/jrione/sso-jrione/domain"
)

type userUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(userRepo domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository: userRepo,
		contextTimeout: timeout,
	}
}

func (u *userUseCase) GetAll(ctx context.Context) (res []domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res, err = u.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (u *userUseCase) GetByUsername(ctx context.Context, username string) (res *domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res, err = u.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return res, nil
}
