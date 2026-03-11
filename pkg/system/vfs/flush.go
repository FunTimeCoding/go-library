package vfs

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
)

func (v *VFS) Flush(dir string) {
	for path, content := range v.files {
		full := filepath.Join(dir, path)
		system.EnsurePathExists(filepath.Dir(full))
		system.SaveFile(full, content)
	}
}
