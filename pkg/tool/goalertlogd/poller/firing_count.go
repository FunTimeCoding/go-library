package poller

func (p *Poller) FiringCount() int {
	return len(p.firing)
}
