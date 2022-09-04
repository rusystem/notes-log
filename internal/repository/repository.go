package repository

import (
	"context"
	"github.com/rusystem/notes-log/internal/config"
	"github.com/rusystem/notes-log/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Insert(ctx context.Context, item domain.LogItem) error
}

type Logger struct {
	cfg *config.Config
	db  *mongo.Database
}

func NewLogRepository(cfg *config.Config, db *mongo.Database) *Logger {
	return &Logger{db: db, cfg: cfg}
}

func (r *Logger) Insert(ctx context.Context, item domain.LogItem) error {
	_, err := r.db.Collection(r.cfg.DB.Collection).InsertOne(ctx, item)

	return err
}
