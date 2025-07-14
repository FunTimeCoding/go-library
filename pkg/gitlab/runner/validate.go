package runner

func (r *Runner) Validate() {
	if r.Paused {
		r.concern = append(r.concern, Paused)
	}
}
