package service

import (
	"context"
	"github.com/rusystem/notes-log/internal/repository"
	"github.com/rusystem/notes-log/pkg/domain"
)

type Service interface {
	Insert(ctx context.Context, req *domain.LogRequest) error
}

type Log struct {
	repo repository.Repository
}

func NewLogService(repo repository.Repository) *Log {
	return &Log{repo: repo}
}

func (s *Log) Insert(ctx context.Context, req *domain.LogRequest) error {
	item := domain.LogItem{
		Action:    req.GetActions().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, item)
}
