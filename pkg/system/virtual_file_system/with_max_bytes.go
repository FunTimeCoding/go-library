package virtual_file_system

func WithMaxBytes(n int64) func(*System) {
	return func(s *System) {
		s.maxBytes = n
	}
}
