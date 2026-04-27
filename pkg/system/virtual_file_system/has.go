package virtual_file_system

func (s *System) Has(path string) bool {
	_, okay := s.files[path]

	return okay
}
