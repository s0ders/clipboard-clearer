package timer

import (
	"time"
)

type ExpirationTimer struct {
	Timer     *time.Timer
	CreatedAt time.Time
	Duration  time.Duration
}

func (c *ExpirationTimer) Stop() {
	c.Timer.Stop()
}

func (c *ExpirationTimer) Update(newDuration time.Duration) {
	// Take into account the time that has already elapsed since the timer started
	newDuration = newDuration - time.Since(c.CreatedAt)

	if newDuration < 0 {
		newDuration = 0
	}

	c.Timer.Reset(newDuration)
}

func New(d time.Duration) *ExpirationTimer {
	return &ExpirationTimer{
		Timer:     time.NewTimer(d),
		CreatedAt: time.Now(),
		Duration:  d,
	}
}
