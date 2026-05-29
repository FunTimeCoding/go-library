package process

import "sync"

func (p *Process) Spawn(
	environment []string,
	waitGroup *sync.WaitGroup,
) {
	p.mutex.Lock()

	if p.handle != nil {
		p.mutex.Unlock()

		return
	}

	if waitGroup != nil {
		waitGroup.Add(1)
	}

	go func() {
		p.run(environment)

		if waitGroup != nil {
			waitGroup.Done()
		}

		p.mutex.Unlock()
	}()
}
