package runner

func (r *Runner) formatShared() string {
	if r.Raw.IsShared {
		return "shared"
	}

	return "not shared"
}
