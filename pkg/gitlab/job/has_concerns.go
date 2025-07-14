package job

func (j *Job) HasConcerns() bool {
	return len(j.concern) > 0
}
