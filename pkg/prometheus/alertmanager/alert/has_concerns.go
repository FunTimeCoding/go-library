package alert

func (a *Alert) HasConcerns() bool {
	return len(a.concern) > 0
}
