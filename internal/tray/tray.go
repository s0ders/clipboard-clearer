// Package tray contains all operations related to the management of the program's system tray icon.
package tray

import (
	"context"
	"fmt"
	"strings"
	"time"

	"fyne.io/systray"

	"github.com/s0ders/clipboard-clearer/internal/appconfig"
)

// Start creates an icon for the program in the system tray.
func Start(ctx context.Context, cancel context.CancelFunc, appConfig *appconfig.Config) {
	onExit := func() {}

	onReady := func() {
		systray.SetTemplateIcon(Data, Data)
		systray.SetTooltip("Clipboard Clearer")

		expirationTimeIndicatorCh := systray.AddMenuItem(FormatDuration(appConfig.ClipboardExpiration()), "")
		systray.AddSeparator()
		increaseCh := systray.AddMenuItem("Increase expiration time", "")
		decreaseCh := systray.AddMenuItem("Decrease expiration time", "")
		systray.AddSeparator()
		quitTrayCh := systray.AddMenuItem("Quit", "Quit the app")

		go func() {
			defer systray.Quit()

			for {
				select {
				case <-increaseCh.ClickedCh:
					appConfig.IncreaseClipboardExpirationTime()
					expirationTimeIndicatorCh.SetTitle(FormatDuration(appConfig.ClipboardExpiration()))
				case <-decreaseCh.ClickedCh:
					appConfig.DecreaseClipboardExpirationTime()
					expirationTimeIndicatorCh.SetTitle(FormatDuration(appConfig.ClipboardExpiration()))
				case <-quitTrayCh.ClickedCh:
					cancel()
					return
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	systray.Run(onReady, onExit)
}

func FormatDuration(d time.Duration) string {
	if d == 0 {
		return "0s"
	}

	var result strings.Builder

	// Define units in descending order
	units := []struct {
		name string
		dur  time.Duration
	}{
		{"h", time.Hour},
		{"m", time.Minute},
		{"s", time.Second},
		{"ms", time.Millisecond},
		{"μs", time.Microsecond},
		{"ns", time.Nanosecond},
	}

	remaining := d

	for _, unit := range units {
		if remaining >= unit.dur {
			count := remaining / unit.dur
			result.WriteString(fmt.Sprintf("%d%s", count, unit.name))
			remaining -= count * unit.dur
		}
	}

	return fmt.Sprintf("Expiration time: %s", result.String())
}
