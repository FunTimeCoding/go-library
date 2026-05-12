package naming

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"path/filepath"
)

func isGenerated(path string) bool {
	return filepath.Base(path) == constant.GeneratedFile
}
