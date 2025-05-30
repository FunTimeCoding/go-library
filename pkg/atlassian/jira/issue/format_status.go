package issue

func (i *Issue) FormatStatus() string {
	if i.Status == "" {
		return NoStatus
	}

	if i.shortStatus != nil {
		return i.shortStatus(i.Status)
	}

	return i.Status
}
