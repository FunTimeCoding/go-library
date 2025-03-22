package option

func WithFormat(format string) func(s *Output) {
	return func(s *Output) {
		s.Format = format
	}
}
