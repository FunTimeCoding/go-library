package runner

func (r *Runner) formatDescription() string {
	if r.Description == "" {
		return NoDescription
	}

	return r.Description
}
