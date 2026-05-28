package process

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/funtimecoding/go-library/pkg/system/writer"
)

func (p *Process) run(
	environment []string,
	errors chan<- error,
) {
	r := run.New()
	r.Panic = false
	r.SetEnvironment(environment)
	r.Writers(p.logger, p.logger)
	r.ProcessGroup()
	handle, e := r.TryOpen("/bin/sh", "-c", p.Command)

	if e != nil {
		select {
		case errors <- e:
		default:
		}

		writer.Print(p.logger, "Failed to start %s: %s\n", p.Name, e)

		return
	}

	p.handle = handle
	p.stoppedBySupervisor = false
	p.mutex.Unlock()
	e = handle.Wait()
	p.mutex.Lock()
	p.condition.Broadcast()

	if e != nil && !p.stoppedBySupervisor {
		select {
		case errors <- e:
		default:
		}
	}

	p.waitError = e
	p.handle = nil
	writer.Print(p.logger, "Terminating %s\n", p.Name)
}
