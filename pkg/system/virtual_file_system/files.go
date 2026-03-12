package virtual_file_system

import "sort"

func (s *System) Files() []string {
	result := make([]string, 0, len(s.files))

	for path := range s.files {
		result = append(result, path)
	}

	sort.Strings(result)

	return result
}
