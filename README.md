# TCP GPS Tracker Simulator (GT06 Protocol)

A high-performance GPS tracker simulator written in Go, designed to stress-test TCP tracking servers using GT06-like binary packets.

--- 

## Overview 

This project emulates GPS tracking devices that communicate over TCP using a GT06-style packet structure.

It allows you to:

- Simulate hundreds or thousands of concurrent tracker devices
- Generate realistic movement patterns
- Introduce configurable latency or jitter
- Test server stability under sustained load
- Validate packet parsing and session handling
- Stress-test connection handling and concurrency behavior

The simulator connects directly to a TCP server and continuously streams tracking packets.

--- 

## Architecture 

Each simulated device: 

1. Opens a TCP connection to the target server
2. Sends login/auth packet (if enabled)
3. Periodically sends location packets
4. Maintains persistent connection
5. Optionally simulates network lag or packet delay

---

## Configuration 

Configuration is done via JSON or CLI flags (for basic configuration). 

Example CLI usage: 
```
go run main.go \
    --ip=127.0.0.1 \ 
    --port=9000 \
    --devices=100 \
    --tick=2s
```

## Running the simulator

1. Clone the repository
```
git clone https://github.com/yourusername/tcp-gps-simulator.git
cd tcp-gps-simulator
```

2. Install dependencies
```
go mod tidy 
``` 

3. Change `config.json` (optional)
  
5. Run the simulator (CLI flags overrides JSON)
```
go run main.go # flags
```

6. Or build a binary

```
go build -o simulator
./simulator
```

---

## Movement simulation modes 

(not available yet)


## License

MIT License 

--- 

## Disclaimer

This project is intended for testing purposes only. 
It does not interact with real GPS Hardware











