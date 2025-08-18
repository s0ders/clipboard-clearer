// Package timer contains operations related to the expiration timer.
package timer

import (
	"time"
)

type ExpirationTimer struct {
	Timer     *time.Timer
	CreatedAt time.Time
	Duration  time.Duration
}

// Stop stops the underlying timer.
func (c *ExpirationTimer) Stop() {
	c.Timer.Stop()
}

// Update replaces the underlying timer with a new one firing at the given new duration. The new duration
// accounts for the time already elapsed since the timer's original creation.
func (c *ExpirationTimer) Update(newDuration time.Duration) {
	if c == nil {
		return
	}

	// Take into account the time that has already elapsed since the timer started
	newDuration = newDuration - time.Since(c.CreatedAt)

	if newDuration < 0 {
		newDuration = 0
	}

	c.Timer.Stop()

	c.Timer = time.NewTimer(newDuration)
}

func New(d time.Duration) *ExpirationTimer {
	return &ExpirationTimer{
		Timer:     time.NewTimer(d),
		CreatedAt: time.Now(),
		Duration:  d,
	}
}
