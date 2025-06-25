package alert

func (a *Alert) Detail(key string) string {
	if result, okay := a.Labels[key]; okay {
		return result
	}

	return ""
}
