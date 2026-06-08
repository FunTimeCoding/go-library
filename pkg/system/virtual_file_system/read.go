package virtual_file_system

func (s *System) Read(path string) []byte {
	f, okay := s.files[path]

	if !okay || !f.Loaded {
		return nil
	}

	return f.Content
}
