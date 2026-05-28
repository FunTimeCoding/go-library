package process

import "io"

func (p *Process) StderrReader() io.ReadCloser {
	return p.stderr
}
