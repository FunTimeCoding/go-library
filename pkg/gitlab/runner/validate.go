package runner

func (r *Runner) validate() []string {
	var result []string

	// TODO: Redundant? its already in the status bubbles
	if false {
		if r.Raw.Online == false {
			result = append(result, "offline")
		}

		if r.Raw.Status != OnlineStatus {
			result = append(result, "not_online")
		}
	}

	if r.Raw.Active == false {
		result = append(result, "inactive")
	}

	if r.Raw.Paused {
		result = append(result, "paused")
	}

	return result
}
