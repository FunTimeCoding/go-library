package run

func Latest(v []*Run) *Run {
	latest := v[0]

	for _, r := range v {
		if r.Status != Completed {
			continue
		}

		if r.CreatedAt.After(latest.CreatedAt) {
			latest = r
		}
	}

	return latest
}
