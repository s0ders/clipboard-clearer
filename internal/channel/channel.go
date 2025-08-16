// Package channel groups operations related to Go channel management.
package channel

import (
	"context"
	"sync"
)

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
