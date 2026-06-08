package virtual_file_system

func (s *System) WriteString(
	path string,
	content string,
) {
	s.Write(path, []byte(content))
}
