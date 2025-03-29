package alert

func (a *Alert) shortenUser(s string) string {
	if a.shortUser != nil {
		return a.shortUser(s)
	}

	return s
}
