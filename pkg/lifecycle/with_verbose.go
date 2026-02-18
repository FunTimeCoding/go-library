package lifecycle

func WithVerbose(v bool) Option {
	return func(l *Lifecycle) {
		l.verbose = v
	}
}
