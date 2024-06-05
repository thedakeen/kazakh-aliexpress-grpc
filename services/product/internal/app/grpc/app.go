package grpcapp

import (
	"fmt"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"net"
	productgrpc "product/internal/grpc/product"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, productService productgrpc.Product, port int) *App {
	gRPCServer := grpc.NewServer()
	productgrpc.Register(gRPCServer, productService, log)

	return &App{
		log:        log,
		port:       port,
		gRPCServer: gRPCServer,
	}
}

func (app *App) MustRun() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}

func (app *App) Run() error {
	const op = "grpcapp.Run"

	log := app.log.With(
		slog.String("op", op),
		slog.Int("port", app.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	err = app.gRPCServer.Serve(l)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (app *App) Stop() {
	const op = "grpcapp.Stop"

	app.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", app.port))

	app.gRPCServer.GracefulStop()
}
