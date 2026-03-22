package process

func (p *Process) Wait() error {
	return p.command.Wait()
}
