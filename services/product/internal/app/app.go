package app

import (
	"golang.org/x/exp/slog"
	grpcapp "product/internal/app/grpc"
	"product/internal/repository"
	"product/internal/services/product"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string) *App {
	storage, err := repository.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := product.New(log, storage, storage)
	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
