package main

import "gt06_sim/internal/sim"

func main() {
	simulator := sim.NewSimulator(sim.DefaultConfig())
	err := simulator.Boot()
	if err != nil {
		panic("Server could not boot")
	}
}
