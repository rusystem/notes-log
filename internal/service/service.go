package service

import (
	"context"
	"github.com/rusystem/notes-log/internal/repository"
	logs "github.com/rusystem/notes-log/pkg/proto"
)

type Logs interface {
	Insert(ctx context.Context, req *logs.LogRequest) (*logs.Empty, error)
}

type Service struct {
	Logs Logs
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Logs: NewLogsService(repo.Logs),
	}
}
