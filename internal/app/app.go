package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/napryag/sso/internal/app/grpc"
	"github.com/napryag/sso/internal/services/auth"
	"github.com/napryag/sso/internal/storage/sqlite"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, ttlToken time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, ttlToken)
	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
