package lifecycle

func (l *Lifecycle) Stop() {
	for i := len(l.component) - 1; i >= 0; i-- {
		l.component[i].Stop(l.verbose)
	}
}
