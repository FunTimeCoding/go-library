package alert

func (a *Alert) SameInstance(o *Alert) bool {
	return a.Instance() == o.Instance()
}
