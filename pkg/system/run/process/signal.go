package process

import "os"

func (p *Process) Signal(s os.Signal) error {
	return p.command.Process.Signal(s)
}
