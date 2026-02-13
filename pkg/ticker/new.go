package ticker

import "time"

func New(
	i time.Duration,
	f func(),
) *Ticker {
	return &Ticker{interval: i, function: f}
}
