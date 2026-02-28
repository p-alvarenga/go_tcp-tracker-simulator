package sim

func (sd *simulatedDevice) startSimulation() {
	for {
		select {
		case <-sd.ctx.Done():
			return
		default:
			return
		}

		//=> Simulation (remove `default: return`)
	}
}
