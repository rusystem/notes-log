package main

import (
	"context"
	"github.com/rusystem/notes-log/internal/config"
	"github.com/rusystem/notes-log/internal/repository"
	"github.com/rusystem/notes-log/internal/server"
	"github.com/rusystem/notes-log/internal/service"
	"github.com/rusystem/notes-log/pkg/database"
	"github.com/sirupsen/logrus"
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

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Ctx.Ttl)
	defer cancel()

	db, err := database.NewMongoConnection(ctx, database.ConnectionInfo{
		URI:      cfg.DB.URI,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		Database: cfg.DB.Database,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	logRepo := repository.NewRepository(cfg, db)
	logService := service.NewService(logRepo)

	srv := server.NewServer(logService.Logs)
	if err := srv.Run(cfg.Server.Port); err != nil {
		logrus.Fatal(err)
	}
	defer srv.Stop()

	logrus.Print("Notes-log started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Notes-log stopped")

}
