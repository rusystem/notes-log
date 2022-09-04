package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionInfo struct {
	URI      string
	Username string
	Password string
	Database string
}

func NewMongoConnection(ctx context.Context, info ConnectionInfo) (*mongo.Database, error) {
	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: info.Username,
		Password: info.Password,
	})
	opts.ApplyURI(info.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	defer func(dbClient *mongo.Client, ctx context.Context) {
		if err := dbClient.Disconnect(ctx); err != nil {
			return
		}
	}(dbClient, ctx)

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return dbClient.Database(info.Database), nil
}
