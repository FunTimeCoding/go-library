package alert

func (a *Alert) nameFromDescription(s string) string {
	if a.descriptionToName != nil {
		return a.descriptionToName(s)
	}

	return s
}
