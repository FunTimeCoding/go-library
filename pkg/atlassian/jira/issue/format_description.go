package issue

func (i *Issue) formatDescription() string {
	if i.Description == "" {
		return NoDescription
	}

	return i.Description
}
