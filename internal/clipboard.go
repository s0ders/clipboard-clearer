package internal

import (
	"context"
	"fmt"
	"time"

	"golang.design/x/clipboard"
)

func WatchAndClearClipboard(ctx context.Context, clipboardExpiration time.Duration) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		if err := clipboard.Init(); err != nil {
			panic(err)
		}

		watchText := clipboard.Watch(ctx, clipboard.FmtText)
		watchImage := clipboard.Watch(ctx, clipboard.FmtImage)

		var contextQueue []context.CancelFunc

		for {
			select {
			case <-ctx.Done():
				return
			case _, ok := <-watchText:
				if !ok {
					return
				}

				if len(contextQueue) > 0 {
					contextQueue[0]()
					contextQueue = contextQueue[:0]
				}

				clearTextContext, clearTextContextFunc := context.WithCancel(ctx)

				contextQueue = append(contextQueue, clearTextContextFunc)

				ClearTextClipboard(clearTextContext, clipboardExpiration)
			case <-watchImage:
				// TODO: not implemented yet
				fmt.Println("New image received")
			}
		}
	}()

	return done
}

func ClearTextClipboard(ctx context.Context, after time.Duration) {
	go func() {
		timer := time.NewTimer(after)

		defer timer.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-timer.C:
				clipboard.Write(clipboard.FmtText, []byte{})
			}
		}
	}()
}
