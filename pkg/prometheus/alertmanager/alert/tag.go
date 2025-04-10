package alert

func (a *Alert) Tag(s ...string) {
	for _, t := range s {
		if !a.HasTag(t) {
			a.Tags = append(a.Tags, t)
		}
	}
}
