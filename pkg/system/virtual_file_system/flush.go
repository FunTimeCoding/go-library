package virtual_file_system

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
)

func (s *System) Flush(dir string) {
	for path, content := range s.files {
		full := filepath.Join(dir, path)
		system.EnsurePathExists(filepath.Dir(full))
		system.SaveFile(full, content)
	}
}
