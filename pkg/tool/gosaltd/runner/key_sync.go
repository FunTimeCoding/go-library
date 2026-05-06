package runner

func (r *Runner) keySync() {
	keys := r.salt.Keys()
	accepted := make(map[string]bool, len(keys.Minions))

	for _, m := range keys.Minions {
		accepted[m] = true
	}

	for _, m := range keys.MinionsDenied {
		if accepted[m] {
			r.logger.Structured("key_delete", "minion", m)
			r.salt.DeleteKey(m)
		}
	}

	for _, m := range keys.MinionsPre {
		r.logger.Structured("key_accept", "minion", m)
		r.salt.AcceptKey(m)
	}
}
