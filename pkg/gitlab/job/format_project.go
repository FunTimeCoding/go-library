package job

func (j *Job) formatProject() string {
	if j.Project != nil {
		return j.Project.CombinedName()
	}

	return NoProject
}
