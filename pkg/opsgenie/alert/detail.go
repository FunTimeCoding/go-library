package alert

func (a *Alert) Detail(k string) string {
	if v, okay := a.Details[k]; okay {
		return v
	}

	return ""
}
