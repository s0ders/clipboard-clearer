package appconfig

import (
	"sync"
	"time"

	"github.com/s0ders/clipboard-clearer/internal/timer"
)

const DefaultExpirationTimesIndex = 2 // 1 minute

var ExpirationTimes = []time.Duration{
	10 * time.Second,
	30 * time.Second,
	1 * time.Minute, // default
	5 * time.Minute,
	10 * time.Minute,
	1 * time.Hour,
}

type Config struct {
	clipboardExpirationTimeIndex int
	mu                           sync.RWMutex
	CurrentTimer                 *timer.ExpirationTimer
}

func (c *Config) ClipboardExpiration() time.Duration {
	defer c.mu.RUnlock()
	c.mu.RLock()

	return ExpirationTimes[c.clipboardExpirationTimeIndex]
}

func (c *Config) DecreaseClipboardExpirationTime() {
	defer c.mu.Unlock()
	c.mu.Lock()

	if c.clipboardExpirationTimeIndex > 0 {
		c.clipboardExpirationTimeIndex--
	}

	c.CurrentTimer.Update(ExpirationTimes[c.clipboardExpirationTimeIndex])
}

func (c *Config) IncreaseClipboardExpirationTime() {
	defer c.mu.Unlock()
	c.mu.Lock()

	if c.clipboardExpirationTimeIndex < len(ExpirationTimes)-1 {
		c.clipboardExpirationTimeIndex++
	}

	c.CurrentTimer.Update(ExpirationTimes[c.clipboardExpirationTimeIndex])
}

func (c *Config) NewExpirationTimer() *timer.ExpirationTimer {
	defer c.mu.Unlock()
	c.mu.Lock()

	if c.CurrentTimer != nil {
		c.CurrentTimer.Stop()
	}

	expiration := ExpirationTimes[c.clipboardExpirationTimeIndex]

	c.CurrentTimer = timer.New(expiration)

	return c.CurrentTimer
}

func New() *Config {
	return &Config{
		clipboardExpirationTimeIndex: DefaultExpirationTimesIndex,
		mu:                           sync.RWMutex{},
	}
}
