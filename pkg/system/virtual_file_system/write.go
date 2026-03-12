package virtual_file_system

func (s *System) Write(
	path string,
	content string,
) {
	s.files[path] = content
}
