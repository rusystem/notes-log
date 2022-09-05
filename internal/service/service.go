package service

import (
	"context"
	"github.com/rusystem/notes-log/internal/repository"
)

type Logs interface {
	Insert(ctx context.Context, message string) error
}

type Service struct {
	Logs Logs
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Logs: NewLogsService(repo.Logs),
	}
}
