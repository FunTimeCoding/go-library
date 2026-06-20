package runner

func (r *Runner) Done() <-chan struct{} {
	return r.stop
}
