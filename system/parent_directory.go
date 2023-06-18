package system

import "path/filepath"

func ParentDirectory(levels int) string {
	var dots []string

	for i := 0; i < levels; i++ {
		dots = append(dots, "..")
	}

	return filepath.Join(append([]string{WorkingDirectory()}, dots...)...)
}
