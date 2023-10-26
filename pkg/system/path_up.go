package system

import "path/filepath"

func PathUp(
	path string,
	levels int,
) string {
	result := path

	for i := 0; i < levels; i++ {
		result = filepath.Dir(result)
	}

	return result
}
