package poller

func (p *Poller) Stop() {
	close(p.stop)
}
