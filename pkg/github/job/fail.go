package job

func (j *Job) Fail() bool {
	return j.Conclusion == Failure
}
