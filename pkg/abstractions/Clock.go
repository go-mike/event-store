package abstractions

import "time"

// Clock is an abstraction for obtaining the current time.
type Clock interface {
	Now() time.Time
}

type systemClock struct{}

// Now implements Clock
func (*systemClock) Now() time.Time {
	return time.Now()
}

// NewSystemClock creates a new system clock.
func NewSystemClock() Clock {
	return &systemClock{}
}
