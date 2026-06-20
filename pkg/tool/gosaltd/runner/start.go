package runner

func (r *Runner) Start() {
	r.provision.Start()
	go r.keySyncLoop()
}
