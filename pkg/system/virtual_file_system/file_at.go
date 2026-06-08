package virtual_file_system

func (s *System) FileAt(path string) *File {
	return s.files[path]
}
