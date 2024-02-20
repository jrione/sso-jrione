package usecase

import (
	"context"
	"time"

	jwtConfig "github.com/jrione/gin-crud/config"
	"github.com/jrione/gin-crud/domain"
)

type loginUseCase struct {
	userRepo domain.UserRepository
	timeout  time.Duration
}

func NewLoginUseCase(ur domain.UserRepository, timeout time.Duration) domain.LoginUseCase {
	return &loginUseCase{
		userRepo: ur,
		timeout:  timeout,
	}
}

func (l *loginUseCase) CheckUser(c context.Context, username string) (loginData *domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, l.timeout)
	defer cancel()

	loginData, err = l.userRepo.GetByUsername(ctx, username)

	return
}

func (l *loginUseCase) CreateAccessToken(user *domain.User, secret string, expire int) (accessToken string, err error) {
	return jwtConfig.CreateAccessToken(user, secret, expire)
}
