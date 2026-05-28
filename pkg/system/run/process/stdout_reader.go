package process

import "io"

func (p *Process) StdoutReader() io.ReadCloser {
	return p.stdout
}
