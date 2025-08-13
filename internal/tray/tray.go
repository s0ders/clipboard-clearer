// Package tray contains all operations related to the management of the program's system tray icon.
package tray

import (
	"context"

	"github.com/getlantern/systray"

	"github.com/s0ders/clipboard-clearer/icon"
)

// Start creates an icon for the program in the system tray.
func Start(ctx context.Context, cancel context.CancelFunc) {
	onExit := func() {}

	onReady := func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		systray.SetTooltip("Clipboard Clearer")

		_ = systray.AddMenuItem("Expiration time: 10s", "")
		systray.AddSeparator()
		_ = systray.AddMenuItem("Increase expiration time", "")
		_ = systray.AddMenuItem("Decrease expiration time", "")
		systray.AddSeparator()
		quitTrayCh := systray.AddMenuItem("Quit", "Quit the app")

		go func() {
			defer systray.Quit()

			for {
				select {
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
