package process

func (p *Process) Running() bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.handle != nil
}
