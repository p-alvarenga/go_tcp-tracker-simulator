package main

import "github.com/p-alvarenga/go_tcp-tracker-simulator/internal/sim"

func main() {
	simulator := sim.NewSimulator(sim.DefaultConfig())
	err := simulator.Boot()
	if err != nil {
		panic("Server could not boot")
	}
}
