package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/napryag/sso/internal/app"
	"github.com/napryag/sso/internal/config"
	"github.com/napryag/sso/internal/lib/logger/handlers/slogpretty"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting app", slog.Any("cfg", cfg))

	appliacation := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go appliacation.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop
	log.Info("stopping application", slog.String("signal", os.Signal.String(sig)))

	appliacation.GRPCSrv.Stop()
	log.Info("application stopped")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	options := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := options.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
