package main

import (
	"context"
	"os"
	"os/signal"
	"product/internal/app"
	"product/internal/config"
	"product/internal/lib/logger/handlers/slogpretty"
	"product/pkg/storage/mongo"
	"syscall"

	"golang.org/x/exp/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	port := 50052

	log := setupLogger(cfg.Env)

	log.Info("starting application",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
		slog.Int("port", port),
	)

	mongo.MustStart(cfg.MongoURI)
	defer func(ctx context.Context) {
		err := mongo.Close(ctx)
		if err != nil {

		}
	}(context.Background())

	application := app.New(log, port, cfg.MongoURI)

	go application.GRPCSrv.MustRun()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	s := <-quit

	log.Info("stopping application", slog.String("signal", s.String()))

	application.GRPCSrv.Stop()

	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
