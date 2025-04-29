package issue

func (i *Issue) FormatLink() string {
	if i.Link == "" {
		return NoLink
	}

	return i.Link
}
