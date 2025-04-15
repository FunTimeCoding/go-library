package issue

func (i *Issue) formatStatus() string {
	if i.Status == "" {
		return NoStatus
	}

	if i.ShortStatus != nil {
		return i.ShortStatus(i.Status)
	}

	return i.Status
}
