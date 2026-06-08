package output

import "strings"

func (r *Results) Relativize(path string) string {
	return strings.TrimPrefix(path, r.workDirectory)
}
