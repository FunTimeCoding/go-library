package relational_tag

func (t *Tag) Nullable() bool {
	return t.nullable && !t.primary
}
