package run

func (r *Run) SetEnvironment(entries []string) {
	r.environment = entries
	r.replaceEnviron = true
}
