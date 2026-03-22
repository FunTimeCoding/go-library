package process

import "os"

func (p *Process) Signal(sig os.Signal) error {
	return p.command.Process.Signal(sig)
}
