package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/spf13/pflag"

	"github.com/dmikhr/sensor-simulator/configs"
	"github.com/dmikhr/sensor-simulator/internal/emitter"
	"github.com/dmikhr/sensor-simulator/internal/sensor"
	"github.com/dmikhr/sensor-simulator/internal/timeutil"
)

const ErrNotSameLength = "sensors config and settings must be of the same length"

func main() {
	var cfg config.Config

	pflag.IntVarP(&cfg.MaxWorkers, "max-workers", "w", 10, "Maximum number of concurrent sensor simulators")
	pflag.IntVarP(&cfg.Port, "port", "p", 8080, "Local server port for the simulation application")
	pflag.DurationVarP(&cfg.EmitterTimeout, "emitter-timeout", "e", 5*time.Second, "Emitter server timeout")
	pflag.StringVarP(&cfg.TargetAddress, "target-address", "a", "localhost", "Target server address to send simulated sensor data")
	pflag.IntVarP(&cfg.TargetPort, "target-port", "t", 9090, "Target server port to send simulated sensor data")
	pflag.StringVarP(&cfg.TargetPath, "target-path", "r", "/api/sensors", "Target server API endpoint path for sensor data")
	pflag.IntVarP(&cfg.SimDuration, "sim-duration", "d", 3, "Simulation duration in seconds")
	pflag.IntVarP(&cfg.ExpCode, "expected-code", "c", http.StatusAccepted, "HTTP response code from target server")
	pflag.BoolVarP(&cfg.DryRun, "dry-run", "u", false, "No actual requests will be sent, only logs will be printed")

	pflag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	sensors, settings := sensor.Init()
	if len(sensors) != len(settings) {
		logger.Error(ErrNotSameLength)
		panic(ErrNotSameLength)
	}

	endpoint := fmt.Sprintf("http://%s:%d/%s", cfg.TargetAddress, cfg.TargetPort, cfg.TargetPath)
	dataSender := emitter.NewHTTPEmitter(endpoint, cfg.EmitterTimeout)
	simulator := emitter.NewSimulator(dataSender, logger, cfg.DryRun)

	ctx, cancel := context.WithTimeout(context.Background(), timeutil.SecToDuration(cfg.SimDuration))
	defer cancel()
	var wg sync.WaitGroup

	enabledSensors := len(sensors)
	if cfg.DryRun {
		slog.Info("Dry run mode enabled")
	}
	slog.Info("Enabled sensors", "count", enabledSensors)
	slog.Info("Starting simulation", "t_sec", cfg.SimDuration)
	for i := 0; i < enabledSensors; i++ {
		slog.Info("Starting simulation", "sensor_id", sensors[i].SensorID,
			"location", settings[i].Location, "frequency_Hz", settings[i].Frequency)
		wg.Add(1)
		go simulator.Simulate(ctx, sensors[i], settings[i].Frequency, &wg)
	}
	wg.Wait()
}
