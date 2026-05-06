package runner

func (r *Runner) connect() {
	r.logger.Structured("salt_connect")
	r.salt = r.saltConnector()
	r.logger.Structured("salt_connected")
}
