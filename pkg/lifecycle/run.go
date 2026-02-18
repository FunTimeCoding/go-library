package lifecycle

func (l *Lifecycle) Run() {
	for _, c := range l.component {
		c.Start()
	}
}
