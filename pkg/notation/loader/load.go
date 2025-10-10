package loader

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func (l *Loader) Load(path string) {
	for _, n := range system.Files(path) {
		l.contents[n] = system.ReadFile(join.Absolute(path, n))
	}
}
