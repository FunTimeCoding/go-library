package process

func (p *Process) Pid() int {
	return p.command.Process.Pid
}
