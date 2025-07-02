package config

import "time"

// Config contains configuration for the application.
type Config struct {
	MaxWorkers     int
	Port           int
	EmitterTimeout time.Duration
	TargetAddress  string
	TargetPort     int
	TargetPath     string
	// in seconds
	SimDuration int
	ExpCode     int
	DryRun      bool
}
