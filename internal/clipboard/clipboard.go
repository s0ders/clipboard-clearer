// Package clipboard contains all operations related to the clipboard management.
package clipboard

import (
	"context"
	"sync"
	"time"

	xclipboard "golang.design/x/clipboard"
)

// WatchAndClear watches the system clipboard and clears it after a given amount of time.
func WatchAndClear(ctx context.Context, clipboardExpiration time.Duration) {
	go func() {
		if err := xclipboard.Init(); err != nil {
			panic(err)
		}

		watchImageChannel := xclipboard.Watch(ctx, xclipboard.FmtImage) // only detects PNG encoded images
		watchTextChannel := xclipboard.Watch(ctx, xclipboard.FmtText)

		watchChannel := FanInChannels(ctx, watchTextChannel, watchImageChannel)

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

				Clear(clearClipboardContext, clipboardExpiration)
			}
		}
	}()
}

// Clear removes the current content of the clipboard.
func Clear(ctx context.Context, after time.Duration) {
	go func() {
		timer := time.NewTimer(after)

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

// FanInChannels sends all messages coming from the input channels into a single output channel.
func FanInChannels[K any](ctx context.Context, channels ...<-chan K) chan K {
	out := make(chan K)

	var wg sync.WaitGroup

	wg.Add(len(channels))

	for _, channel := range channels {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case v, more := <-channel:
					if !more {
						return
					}

					select {
					case <-ctx.Done():
						return
					case out <- v:
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
