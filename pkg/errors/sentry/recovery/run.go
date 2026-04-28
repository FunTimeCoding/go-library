package recovery

func (r *Recovery) Run(f func()) {
	defer func() {
		if v := recover(); v != nil {
			if r.hub != nil {
				r.hub.Recover(v)
			}

			r.logger.Plain("recovered from panic: %v", v)
		}
	}()
	f()
}
