package virtual_file_system

import (
	"fmt"
	"strings"
)

func (s *System) DirectoryExists(path string) bool {
	prefix := fmt.Sprintf("%s/", path)

	for key := range s.files {
		if strings.HasPrefix(key, prefix) {
			return true
		}
	}

	return false
}
