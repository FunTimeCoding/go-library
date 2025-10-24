package amtool

import (
	"github.com/funtimecoding/go-library/pkg/markup"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Read(
	base string,
	name string,
) *Option {
	c := &Option{}
	content := system.ReadFile(base, name)
	markup.DecodeStrict(content, &c)

	return c
}
