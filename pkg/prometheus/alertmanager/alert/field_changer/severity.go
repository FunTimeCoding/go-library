package field_changer

func (a *Changer) Severity(
	alert string,
	name string,
) string {
	for _, s := range a.severity {
		if s.Match(alert, name) {
			return s.To
		}
	}

	return name
}
