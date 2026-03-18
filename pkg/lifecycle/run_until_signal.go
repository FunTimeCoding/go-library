package lifecycle

import "github.com/funtimecoding/go-library/pkg/system"

func (l *Lifecycle) RunUntilSignal() {
	l.Run()
	system.KillSignalBlock()
	l.Stop()
}
