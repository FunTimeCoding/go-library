package web

import "path/filepath"

func seedDirectory(path string) string {
	d := filepath.Dir(path)

	if d == "." {
		return ""
	}

	return d
}
