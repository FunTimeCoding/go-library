package runner

func (r *Runner) Stop() {
	close(r.stop)
}
