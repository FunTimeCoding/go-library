package issue

func (i *Issue) HasConcerns() bool {
	return len(i.concerns) > 0
}
