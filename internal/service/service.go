package service

import (
	"context"
	"github.com/rusystem/notes-log/internal/repository"
	"github.com/rusystem/notes-log/pkg/domain"
)

type Logs interface {
	Insert(ctx context.Context, req *domain.LogRequest) (*domain.Empty, error)
}

type Service struct {
	Logs Logs
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Logs: NewLogsService(repo.Logs),
	}
}
