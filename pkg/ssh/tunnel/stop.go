package tunnel

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"syscall"
)

func (t *Tunnel) Stop() {
	if t.command == nil {
		panic("not running")
	}

	e := t.command.Process.Signal(syscall.SIGTERM)

	if e != nil && e.Error() != "os: process already finished" {
		errors.PanicOnError(e)
	}

	<-t.stopped
}
