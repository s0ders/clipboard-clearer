// Package tray contains all operations related to the management of the program's system tray icon.
package tray

import (
	"context"

	"github.com/getlantern/systray"

	"github.com/s0ders/clipboard-clearer/icon"
)

// StartSystray creates an icon for the program in the system tray.
func StartSystray(ctx context.Context, cancel context.CancelFunc) {
	go func() {
		onExit := func() {}

		onReady := func() {
			systray.SetTemplateIcon(icon.Data, icon.Data)
			systray.SetTitle("Clipboard Clearer")
			systray.SetTooltip("Clipboard Clearer")
			quitTrayCh := systray.AddMenuItem("Quit", "Quit the app")

			go func() {
				for {
					select {
					case <-quitTrayCh.ClickedCh:
						cancel()
						systray.Quit()
					case <-ctx.Done():
						systray.Quit()
					}
				}
			}()
		}

		systray.Run(onReady, onExit)
	}()
}
