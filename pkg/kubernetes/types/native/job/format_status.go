package job

func (j *Job) formatStatus() string {
	if j.Raw.Status.Active > 0 {
		return "active"
	}

	if j.Raw.Status.Succeeded > 0 && j.Raw.Status.Failed > 0 {
		return "mixed"
	}

	if j.Raw.Status.Succeeded > 0 {
		return "succeeded"
	}

	if j.Raw.Status.Failed > 0 {
		return "failed"
	}

	return "unknown"
}
