package runner

func (r *Runner) formatShared() string {
	if r.Shared {
		return "shared"
	}

	return "not shared"
}
