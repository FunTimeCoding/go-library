package runner

func (r *Runner) HasConcerns() bool {
	return len(r.concern) > 0
}
