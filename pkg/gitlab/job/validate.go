package job

func (j *Job) Validate() {
	if j.Fail() {
		j.concern = append(j.concern, Failed)
	}
}
