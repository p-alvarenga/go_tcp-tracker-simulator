package main

import (
	"go_tcp-tracker-simulator/internal/config"
	"go_tcp-tracker-simulator/internal/simulator"
	"log/slog"
)

func main() {
	config := config.DefaultConfig()

	sim := simulator.NewSimulator(*config, slog.Default())
	sim.Boot()
}
