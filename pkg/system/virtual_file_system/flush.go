package virtual_file_system

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
)

func (s *System) Flush(directory string) {
	for path, content := range s.files {
		full := filepath.Join(directory, path)
		system.EnsurePathExists(filepath.Dir(full))
		system.SaveFile(full, content)
	}
}
