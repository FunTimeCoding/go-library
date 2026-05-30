package server

import "github.com/funtimecoding/go-library/pkg/identity"

func New(
	i *identity.Tool,
	version string,
) *Builder {
	return &Builder{
		name:         i.Name(),
		version:      version,
		instructions: i.Instructions(),
	}
}
