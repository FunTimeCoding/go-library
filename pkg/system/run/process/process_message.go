package process

func (p *Process) ProcessMessage() string {
	return p.command.ProcessState.String()
}
