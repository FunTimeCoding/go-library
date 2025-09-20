package cron_job

func (j *Job) formatStatus() string {
	if len(j.Raw.Status.Active) > 0 {
		return "active"
	}

	return "inactive"
}
