package option

func WithDebug() func(s *Output) {
	return func(s *Output) {
		s.Debug = true
	}
}
