package config

import (
	"flag"
	"fmt"
	"os"
)

func ParseFlags(cfg *SimulatorConfig) {
	flag.StringVar(&cfg.ServerHost, "host", cfg.ServerHost, "server host (e.g. 0.0.0.0, 127.0.0.1)")
	flag.IntVar(&cfg.ServerPort, "port", cfg.ServerPort, "server port")
	flag.IntVar(&cfg.MaxDevices, "devices", cfg.MaxDevices, "number of devices simulated")
	flag.DurationVar(&cfg.SimulatedDeviceConfig.TickInterval, "tick", cfg.SimulatedDeviceConfig.TickInterval, "tick interval (e.g. 500ms, 2s, 1m)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "GO TCP Tracker Simulator\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  go_tcp-tracker-simulator [options]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()
}
