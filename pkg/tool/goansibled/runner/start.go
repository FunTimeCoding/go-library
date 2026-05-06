package runner

func (r *Runner) Start() {
	go r.run()
}
