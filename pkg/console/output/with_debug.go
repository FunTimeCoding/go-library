package output

func WithDebug() func(s *Settings) {
	return func(s *Settings) {
		s.Debug()
	}
}
