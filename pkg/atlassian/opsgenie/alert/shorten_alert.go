package alert

func (a *Alert) shortenAlert(s string) string {
	if a.shortAlert != nil {
		return a.shortAlert(s)
	}

	return s
}
