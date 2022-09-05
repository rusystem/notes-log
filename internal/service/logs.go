package service

import (
	"context"
	"encoding/json"
	"github.com/rusystem/notes-log/internal/domain"
	"github.com/rusystem/notes-log/internal/repository"
)

type LogsService struct {
	repo repository.Logs
}

func NewLogsService(repo repository.Logs) *LogsService {
	return &LogsService{
		repo: repo,
	}
}

func (s *LogsService) Insert(ctx context.Context, message string) error {
	var data domain.LogItem
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		return err
	}

	item := domain.LogItem{
		Action:    data.Action,
		Entity:    data.Entity,
		EntityID:  data.EntityID,
		Timestamp: data.Timestamp,
	}

	return s.repo.Insert(ctx, item)
}
