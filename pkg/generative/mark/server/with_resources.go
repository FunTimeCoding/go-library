package server

func (b *Builder) WithResources() *Builder {
	b.resources = true

	return b
}
