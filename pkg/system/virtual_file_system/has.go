package virtual_file_system

func (s *System) Has(path string) bool {
	_, ok := s.files[path]

	return ok
}
