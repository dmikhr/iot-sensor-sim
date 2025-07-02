package sensor

import (
	"fmt"
	"math/rand"
	"time"
)

type ValueType string

const (
	ValueTypePressure    ValueType = "pressure"
	ValueTypeTemperature ValueType = "temperature"
	ValueTypeVoltage     ValueType = "voltage"
)

// Data contains configuration information for sensor simulations.
type Data struct {
	SensorID  string
	Parameter ValueType
	Unit      string
	Frequency float64
	Location  string
	Enabled   bool
}

// Sensor represents data emitted by sensor.
type Sensor struct {
	SensorID  string    `json:"sensorId"`
	Parameter ValueType `json:"type"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

// New creates new sensor instance.
func New(id string, param ValueType, unit string) *Sensor {
	return &Sensor{
		SensorID:  id,
		Parameter: param,
		Unit:      unit,
	}
}

// Settings contains sensor settings.
type Settings struct {
	SensorID  string
	Frequency float64
	Location  string
}

// NewSettings creates a new settings instance.
func NewSettings(id string, freq float64, location string) *Settings {
	return &Settings{
		SensorID: id,
		// Hz
		Frequency: freq,
		Location:  location,
	}
}

// GenerateReading generates new sensor reading.
func (s *Sensor) GenerateReading() {
	s.Timestamp = time.Now()

	switch s.Parameter {
	case ValueTypePressure:
		s.Value = newP()
	case ValueTypeTemperature:
		s.Value = newT()
	case ValueTypeVoltage:
		s.Value = newV()
	}
}

func newT() float64 {
	return 25
}

func newP() float64 {
	return 100000
}

func newV() float64 {
	return 400
}

// rndRange returns a random number in the range [min, max].
func rndRange(r *rand.Rand, minVal, maxVal float64) float64 {
	if minVal > maxVal {
		panic(fmt.Sprintf("invalid range: min (%.6f) > max (%.6f)", minVal, maxVal))
	}
	return minVal + r.Float64()*(maxVal-minVal)
}
