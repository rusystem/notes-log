package main

import (
	"context"
	"github.com/rusystem/notes-log/internal/config"
	"github.com/rusystem/notes-log/internal/repository"
	"github.com/rusystem/notes-log/internal/server"
	"github.com/rusystem/notes-log/internal/service"
	"github.com/rusystem/notes-log/pkg/database"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"os/signal"
	"syscall"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.ErrorLevel)
}

func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		logrus.Fatal(err)
	}

	ctx := context.Background()

	dbClient, err := database.NewMongoClient(ctx, database.ConnectionInfo{
		URI:      cfg.DB.URI,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	defer func(dbClient *mongo.Client, ctx context.Context) {
		if err := dbClient.Disconnect(ctx); err != nil {
			logrus.Fatal(err)
		}
	}(dbClient, ctx)
	db := dbClient.Database(cfg.DB.Database)

	logRepo := repository.NewRepository(cfg, db)
	logService := service.NewService(logRepo)

	srv := server.NewServer(logService)
	if err := srv.Listen(cfg); err != nil {
		logrus.Fatal(err)
	}
	defer func(srv *server.Server) {
		if err := srv.Close(); err != nil {
			logrus.Fatal(err)
		}
	}(srv)

	go func() {
		if err := srv.Serve(ctx); err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Info("Notes-log started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("Notes-log stopped")

}
