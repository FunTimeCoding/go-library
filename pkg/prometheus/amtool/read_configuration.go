package amtool

import (
	"github.com/funtimecoding/go-library/pkg/markup"
	"github.com/funtimecoding/go-library/pkg/system"
)

func ReadConfiguration(path string) *Configuration {
	c := &Configuration{}
	content := system.ReadFile(path)
	markup.DecodeStrict(content, &c)

	return c
}
