package runner

func (r *Runner) formatName() string {
	if r.Name == "" {
		return NoName
	}

	return r.Name
}
