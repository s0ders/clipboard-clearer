package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/getlantern/systray"

	"github.com/s0ders/clipboard-clearer/icon"
	"github.com/s0ders/clipboard-clearer/internal"
)

func main() {
	var interruptSignal = make(chan os.Signal, 1)
	signal.Notify(interruptSignal, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-interruptSignal
		cancel()
	}()

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

	done := internal.WatchAndClearClipboard(ctx, 5*time.Second)

	<-done
}
