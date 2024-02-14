package domain

import "context"

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json"age"`
	Grade int    `json"grade"`
}

type StudentUseCase interface {
	Fetch(ctx context.Context) ([]Student, error)
}

type StudentRepository interface {
	Fetch(ctx context.Context) (res []Student, err error)
}
