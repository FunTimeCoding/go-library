package runner

func (r *Runner) validate() []string {
	var result []string

	// TODO: Redundant? its already in the status bubbles
	if false {
		if !r.Online {
			result = append(result, "offline")
		}

		if r.Status != OnlineStatus {
			result = append(result, "not_online")
		}
	}

	if r.Paused {
		result = append(result, "paused")
	}

	return result
}
