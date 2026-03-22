package process

func (p *Process) ExitCode() int {
	return p.command.ProcessState.ExitCode()
}
