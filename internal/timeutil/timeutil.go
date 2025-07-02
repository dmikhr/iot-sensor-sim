// Package timeutil provides functions for working with time.
package timeutil

import "time"

// SecToDuration converts seconds to time.Duration for timeout context.
func SecToDuration(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}

// HzToDuration convert Hz to time.Duration.
func HzToDuration(hz float64) time.Duration {
	if hz <= 0 {
		panic("Hz must be positive")
	}
	seconds := 1.0 / hz
	return time.Duration(seconds * float64(time.Second))
}
