package virtual_file_system

import "time"

func (s *System) Write(
	path string,
	content []byte,
) {
	s.checkLimits(len(content))
	s.files[path] = &File{
		Content: content,
		Size:    int64(len(content)),
		ModTime: time.Now(),
		Loaded:  true,
	}
	s.written[path] = true
}
