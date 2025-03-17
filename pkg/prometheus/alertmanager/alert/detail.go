package alert

func (a *Alert) Detail(key string) string {
	if result, okay := a.Raw.Labels[key]; okay {
		return result
	}

	return ""
}
