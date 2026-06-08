package virtual_file_system

func (s *System) ReadString(path string) string {
	return string(s.Read(path))
}
