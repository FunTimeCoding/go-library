package process

func (p *Process) Log() []string {
	return p.logger.History()
}
