package poller

import "time"

func (p *Poller) Start() {
	go func() {
		t := time.NewTicker(p.interval)
		defer t.Stop()
		p.recovery.Run(p.Poll)

		for {
			select {
			case <-t.C:
				p.recovery.Run(p.Poll)
			case <-p.stop:
				return
			}
		}
	}()
}
