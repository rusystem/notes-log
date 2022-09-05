package service

import (
	"context"
	"github.com/rusystem/notes-log/internal/repository"
	"github.com/rusystem/notes-log/pkg/domain"
	logs "github.com/rusystem/notes-log/pkg/proto"
)

type LogsService struct {
	repo repository.Logs
}

func NewLogsService(repo repository.Logs) *LogsService {
	return &LogsService{
		repo: repo,
	}
}

func (s *LogsService) Insert(ctx context.Context, req *logs.LogRequest) (*logs.Empty, error) {
	item := domain.LogItem{
		Action:    req.GetActions().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return &logs.Empty{}, s.repo.Insert(ctx, item)
}
