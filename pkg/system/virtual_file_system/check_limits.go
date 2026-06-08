package virtual_file_system

import "fmt"

func (s *System) checkLimits(additional int) {
	if s.maxFiles > 0 && len(s.files) >= s.maxFiles {
		panic(
			fmt.Sprintf(
				"virtual file system: max files exceeded (%d)",
				s.maxFiles,
			),
		)
	}

	if s.maxBytes > 0 && s.totalBytes()+int64(additional) > s.maxBytes {
		panic(
			fmt.Sprintf(
				"virtual file system: max bytes exceeded (%d)",
				s.maxBytes,
			),
		)
	}
}
