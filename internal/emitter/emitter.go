// Package emitter provides functions for sending http request simulating smart sensors activity.
package emitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dmikhr/sensor-simulator/internal/sensor"
)

// Emitter defines an interface for types that can emit sensor data.
type Emitter interface {
	Emit(sensor.Sensor) error
}

// HTTPEmitter defines struct for sending sensor data via http to a given endpoint.
type HTTPEmitter struct {
	client   *http.Client
	endpoint string
}

// NewHTTPEmitter creates a new emitter instance.
func NewHTTPEmitter(endpoint string, timeout time.Duration) *HTTPEmitter {
	return &HTTPEmitter{
		client:   &http.Client{Timeout: timeout},
		endpoint: endpoint,
	}
}

// Emit sends sensor data to the endpoint.
func (e *HTTPEmitter) Emit(s sensor.Sensor) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, e.endpoint, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := e.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 202 - data accepted but not yet been processed
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("HTTP error: %d %s", resp.StatusCode, resp.Status)
	}

	return nil
}
