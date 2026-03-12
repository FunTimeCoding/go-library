package virtual_file_system

func (s *System) Read(path string) string {
	return s.files[path]
}
