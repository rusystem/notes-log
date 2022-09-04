package service

import (
	"context"
	"github.com/rusystem/notes-log/internal/repository"
	"github.com/rusystem/notes-log/pkg/domain"
	log "github.com/rusystem/notes-log/pkg/domain"
)

type LogsService struct {
	repo repository.Logs
	log.UnimplementedLogsServer
}

func NewLogsService(repo repository.Logs) *LogsService {
	return &LogsService{
		repo: repo,
	}
}

func (s *LogsService) Insert(ctx context.Context, req *domain.LogRequest) (*domain.Empty, error) {
	item := domain.LogItem{
		Action:    req.GetActions().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return &domain.Empty{}, s.repo.Insert(ctx, item)
}
