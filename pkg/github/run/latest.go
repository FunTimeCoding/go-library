package run

func Latest(v []*Run) *Run {
	var latest *Run

	for _, r := range v {
		if r.Status != Completed {
			continue
		}

		if latest == nil {
			latest = r

			continue
		}

		if r.CreatedAt.After(latest.CreatedAt) {
			latest = r
		}
	}

	return latest
}
