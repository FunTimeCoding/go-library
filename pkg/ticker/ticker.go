package ticker

import "time"

type Ticker struct {
	interval time.Duration
	function func()
	ticker   *time.Ticker
	done     chan struct{}
}
