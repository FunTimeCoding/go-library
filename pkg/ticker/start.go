package ticker

import "time"

func (t *Ticker) Start() {
	t.ticker = time.NewTicker(t.interval)
	t.done = make(chan struct{})

	go func() {
		for {
			select {
			case <-t.done:
				return
			case <-t.ticker.C:
				t.function()
			}
		}
	}()
}
