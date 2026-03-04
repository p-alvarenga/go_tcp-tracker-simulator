package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/config"
	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/sim"
)

func main() {
	logger := buildLogger()

	cfg := config.DefaultConfig()
	config.ParseFlags(cfg)

	simulator := sim.NewSimulator(cfg, logger)
	_ = simulator.Boot()
}

func buildLogger() *slog.Logger {
	if os.Getenv("APP_ENV") == "production" {
		return slog.New(
			slog.NewJSONHandler(os.Stdout, nil),
		)
	}

	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				switch a.Key {
				case slog.TimeKey:
					return slog.String("t", a.Value.Time().Format(time.RFC3339))
				case slog.LevelKey:
					return slog.String("lvl", a.Value.String())
				case slog.SourceKey:
					source := a.Value.Any().(*slog.Source)
					source.File = filepath.Base(source.File)
					return slog.Attr{
						Key:   slog.SourceKey,
						Value: slog.AnyValue(source),
					}

				default:
					return a
				}
			},
		}),
	)
}
