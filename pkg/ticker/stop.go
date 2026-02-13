package ticker

func (t *Ticker) Stop() {
	t.ticker.Stop()
	close(t.done)
}
