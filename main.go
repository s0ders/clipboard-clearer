package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/s0ders/clipboard-clearer/internal/appconfig"
	"github.com/s0ders/clipboard-clearer/internal/clipboard"
	"github.com/s0ders/clipboard-clearer/internal/tray"
)

func main() {
	appConfig := appconfig.New()

	var interruptSignal = make(chan os.Signal, 1)
	signal.Notify(interruptSignal, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-interruptSignal
		cancel()
	}()

	clipboard.WatchAndClear(ctx, appConfig)

	tray.Start(ctx, cancel, appConfig)
}
