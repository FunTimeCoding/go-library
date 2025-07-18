package job

func (j *Job) ProjectLink() string {
	if j.Project != nil {
		return j.Project.Link
	}

	return ""
}
