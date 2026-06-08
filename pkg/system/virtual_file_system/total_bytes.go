package virtual_file_system

func (s *System) totalBytes() int64 {
	var total int64

	for _, f := range s.files {
		if f.Loaded {
			total += int64(len(f.Content))
		}
	}

	return total
}
