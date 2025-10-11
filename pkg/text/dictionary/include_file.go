package dictionary

import (
	"github.com/funtimecoding/go-library/pkg/text/dictionary/constant"
	"path/filepath"
	"strings"
)

func IncludeFile(name string) bool {
	e := filepath.Ext(name)

	if e == "" {
		return constant.NoExtension[name]
	}

	return constant.Extensions[strings.ToLower(e)]
}
