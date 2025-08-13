package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/s0ders/clipboard-clearer/internal/clipboard"
	"github.com/s0ders/clipboard-clearer/internal/tray"
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

	clipboard.WatchAndClear(ctx, 5*time.Second)

	tray.Start(ctx, cancel)
}
