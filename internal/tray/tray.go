// Package tray contains all operations related to the management of the program's system tray icon.
package tray

import (
	"context"
	"fmt"
	"time"

	"github.com/getlantern/systray"

	"github.com/s0ders/clipboard-clearer/icon"
	"github.com/s0ders/clipboard-clearer/internal/appconfig"
)

// Start creates an icon for the program in the system tray.
func Start(ctx context.Context, cancel context.CancelFunc, appConfig *appconfig.Config) {
	onExit := func() {}

	onReady := func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
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
	durationString := d.String()

	return fmt.Sprintf("Expiration time: %s", durationString)
}
