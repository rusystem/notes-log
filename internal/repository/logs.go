package repository

import (
	"context"
	"github.com/rusystem/notes-log/internal/config"
	"github.com/rusystem/notes-log/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogsRepository struct {
	cfg *config.Config
	db  *mongo.Database
}

func NewLogsRepository(cfg *config.Config, db *mongo.Database) *LogsRepository {
	return &LogsRepository{
		cfg: cfg,
		db:  db,
	}
}

func (r *LogsRepository) Insert(ctx context.Context, item domain.LogItem) error {
	_, err := r.db.Collection(r.cfg.DB.Collection).InsertOne(ctx, item)

	return err
}
