package output

func WithFormat(format string) func(s *Settings) {
	return func(s *Settings) {
		s.Format = format
	}
}
