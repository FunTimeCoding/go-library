package amtool

import (
	"github.com/funtimecoding/go-library/pkg/markup"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Read(path string) *Option {
	c := &Option{}
	content := system.ReadFile(path)
	markup.DecodeStrict(content, &c)

	return c
}
