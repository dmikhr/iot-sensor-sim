package emitter

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/dmikhr/sensor-simulator/internal/sensor"
	"github.com/dmikhr/sensor-simulator/internal/timeutil"
)

// var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Simulator struct {
	sender *HTTPEmitter
	logger *slog.Logger
	dryRun bool
}

func NewSimulator(sender *HTTPEmitter, logger *slog.Logger, dryRun bool) *Simulator {
	return &Simulator{
		sender: sender,
		logger: logger,
		dryRun: dryRun,
	}
}

// Simulate - simulate sensor emitting data
func (s *Simulator) Simulate(ctx context.Context, sensor *sensor.Sensor, frequency float64, wg *sync.WaitGroup) {
	defer wg.Done()

	interval := timeutil.HzToDuration(frequency)
	s.logger.Info("Setting emitting interval", "sensor_id", sensor.SensorID,
		"interval_sec", interval, "frequency_Hz", frequency)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			sensor.GenerateReading()
			if s.dryRun {
				s.logger.Info("Dry run mode enabled, skipping sending sensor data", "sensor_id", sensor.SensorID)
			} else {
				err := s.sender.Emit(*sensor)
				if err != nil {
					s.logger.Error("error sending sensor data", "error", err)
				} else {
					s.logger.Info("sent sensor data", "sensor_id", sensor.SensorID)
				}
			}
		case <-ctx.Done():
			s.logger.Info("Finished simulation", "sensor_id", sensor.SensorID)
			return
		}
	}
}
