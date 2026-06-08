package virtual_file_system

func (s *System) Delete(path string) {
	delete(s.files, path)
	s.deleted[path] = true
}
