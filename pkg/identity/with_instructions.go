package identity

func (t *Tool) WithInstructions(v string) *Tool {
	t.instructions = v

	return t
}
