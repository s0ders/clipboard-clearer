// Package clipboard contains all operations related to the clipboard management.
package clipboard

import (
	"context"
	"time"

	xclipboard "golang.design/x/clipboard"

	"github.com/s0ders/clipboard-clearer/internal/appconfig"
	"github.com/s0ders/clipboard-clearer/internal/channel"
)

// WatchAndClear watches the system clipboard and clears it after a given amount of time.
func WatchAndClear(ctx context.Context, appConfig *appconfig.Config) {
	go func() {
		if err := xclipboard.Init(); err != nil {
			panic(err)
		}

		watchImageChannel := xclipboard.Watch(ctx, xclipboard.FmtImage) // only detects PNG encoded images
		watchTextChannel := xclipboard.Watch(ctx, xclipboard.FmtText)

		watchChannel := channel.FanInChannels(ctx, watchTextChannel, watchImageChannel)

		var contextQueue []context.CancelFunc

		for {
			select {
			case <-ctx.Done():
				return
			case _, ok := <-watchChannel:
				if !ok {
					return
				}

				if len(contextQueue) > 0 {
					contextQueue[0]()
					contextQueue = contextQueue[:0]
				}

				clearClipboardContext, clearClipboardContextFunc := context.WithCancel(ctx)

				contextQueue = append(contextQueue, clearClipboardContextFunc)

				Clear(clearClipboardContext, appConfig)
			}
		}
	}()
}

// TODO: see above TODO, would need to stop existing timer and restart them with the new
// delay.

// Clear removes the current content of the clipboard.
func Clear(ctx context.Context, appConfig *appconfig.Config) {
	go func() {
		timer := time.NewTimer(appConfig.ClipboardExpiration())

		defer timer.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-timer.C:
				xclipboard.Write(xclipboard.FmtText, []byte{})
				return
			}
		}
	}()
}
