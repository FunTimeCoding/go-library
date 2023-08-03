package system

import "path/filepath"

func Join(s ...string) string {
	return filepath.Join(s...)
}
