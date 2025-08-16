package appconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppConfig_UpdateClipboardExpirationTime(t *testing.T) {
	cfg := New()

	assert.Equal(t, ExpirationTimes[DefaultExpirationTimesIndex], cfg.ClipboardExpiration(), "incorrect default expiration time")

	cfg.IncreaseClipboardExpirationTime()

	assert.Equal(t, ExpirationTimes[DefaultExpirationTimesIndex+1], cfg.ClipboardExpiration(), "incorrect expiration time")

	cfg.DecreaseClipboardExpirationTime()

	assert.Equal(t, ExpirationTimes[DefaultExpirationTimesIndex], cfg.ClipboardExpiration(), "incorrect expiration time")
}
