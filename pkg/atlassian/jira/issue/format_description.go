package issue

func (i *Issue) FormatDescription() string {
	if i.Description == "" {
		return NoDescription
	}

	return i.Description
}
