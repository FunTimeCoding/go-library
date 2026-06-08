package virtual_file_system

func WithMaxFiles(n int) func(*System) {
	return func(s *System) {
		s.maxFiles = n
	}
}
