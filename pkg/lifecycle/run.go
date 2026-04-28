package lifecycle

func (l *Lifecycle) Run() {
	for _, c := range l.component {
		c.Start()
	}

	if l.logger != nil {
		l.logger.Structured("lifecycle_start")
	}
}
