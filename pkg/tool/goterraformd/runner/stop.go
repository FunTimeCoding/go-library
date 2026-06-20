package runner

func (r *Runner) Stop() {
	r.provision.Stop()
}
