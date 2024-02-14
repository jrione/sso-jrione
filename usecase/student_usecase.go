package usecase

import (
	"context"
	"time"

	"github.com/jrione/gin-crud/domain"
)

type studentUseCase struct {
	studentRepo    domain.StudentRepository
	contextTimeout time.Duration
}

func NewStudentUseCase(StudentRepository domain.StudentRepository, timeout time.Duration) domain.StudentUseCase {
	return &studentUseCase{
		studentRepo:    StudentRepository,
		contextTimeout: timeout,
	}
}

func (a *studentUseCase) Fetch(c context.Context) (res []domain.Student, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.studentRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}
