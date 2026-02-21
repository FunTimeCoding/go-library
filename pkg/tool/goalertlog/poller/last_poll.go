package poller

import "time"

func (p *Poller) LastPoll() time.Time {
	v := p.lastPoll.Load()

	if v == nil {
		return time.Time{}
	}

	return v.(time.Time)
}
