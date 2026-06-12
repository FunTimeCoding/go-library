package server

func (b *Builder) WithInstructions(v string) *Builder {
	b.instructions = v

	return b
}
